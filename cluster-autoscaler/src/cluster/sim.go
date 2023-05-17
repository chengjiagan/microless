package cluster

import (
	"context"
	"fmt"
	"microless/cluster-autoscaler/spec"
	"sort"
	"time"

	corev1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1"
	optsv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	Manager = "sim-cluster-manager"
)

type simClusterManager struct {
	// params from config file
	serverlessCpu      float64
	serverlessMem      float64
	vmCpu              float64
	vmMem              float64
	serverlessCpuPrice float64
	serverlessMemPrice float64
	vmPrice            float64
	scaleUpLatency     time.Duration
	namespace          string

	c *kubernetes.Clientset
}

// simulating adding a new node
// actually just set the node to schedulable
// TODO: simulate the time to add a new node
func (cm *simClusterManager) ScaleUpNode(n int) error {
	// get free nodes
	ctx := context.Background()
	opts := optsv1.ListOptions{
		FieldSelector: "spec.unschedulable=true",
		LabelSelector: "type=vm",
	}
	nodes, err := cm.c.CoreV1().Nodes().List(ctx, opts)
	if err != nil {
		return fmt.Errorf("failed to list nodes: %v", err)
	}

	if len(nodes.Items) < n {
		klog.Warningf("not enough nodes (%d) to scale up", len(nodes.Items))
		n = len(nodes.Items)
	}

	for i := 0; i < n; i++ {
		// patch node to make it schedulable
		nodeName := nodes.Items[i].Name
		_, err := cm.c.CoreV1().Nodes().Patch(
			ctx,
			nodeName,
			types.MergePatchType,
			[]byte(`{"spec":{"unschedulable":false}}`),
			optsv1.PatchOptions{FieldManager: Manager},
		)
		if err != nil {
			return fmt.Errorf("failed to update node %s: %v", nodeName, err)
		}
	}
	return nil
}

type nodeInfo struct {
	name string
	pods []string
}
type nodeSlice []nodeInfo

func (ns nodeSlice) Len() int {
	return len(ns)
}

func (ns nodeSlice) Less(i, j int) bool {
	return len(ns[i].pods) < len(ns[j].pods)
}

func (ns nodeSlice) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

// simulating removing a node
// actually just set the node to unschedulable and evict all pods on it
func (cm *simClusterManager) ScaleDownNode(n int) error {
	// get busy nodes
	ctx := context.Background()
	nodeClient := cm.c.CoreV1().Nodes()
	nodes, err := nodeClient.List(
		ctx,
		optsv1.ListOptions{
			FieldSelector: "spec.unschedulable=false",
			LabelSelector: "type=vm",
		},
	)
	if err != nil {
		return fmt.Errorf("failed to list nodes: %v", err)
	}

	if len(nodes.Items) < n {
		return fmt.Errorf("not enough nodes to scale down")
	}

	// sort nodes by number of pods on it
	podClient := cm.c.CoreV1().Pods(cm.namespace)
	ns := make(nodeSlice, len(nodes.Items))
	for i := range ns {
		name := nodes.Items[i].Name

		// get pods on the node
		pods, err := podClient.List(
			ctx,
			optsv1.ListOptions{FieldSelector: "spec.nodeName=" + name},
		)
		if err != nil {
			return fmt.Errorf("failed to list pods on node %s: %v", name, err)
		}

		// get pod names
		podNames := make([]string, len(pods.Items))
		for j := range podNames {
			podNames[j] = pods.Items[j].Name
		}

		ns[i] = nodeInfo{name, podNames}
	}
	sort.Sort(ns)

	evictClient := cm.c.PolicyV1().Evictions(cm.namespace)
	for i := 0; i < n; i++ {
		// patch node as unchedulable
		_, err := nodeClient.Patch(
			ctx,
			ns[i].name,
			types.MergePatchType,
			[]byte(`{"spec":{"unschedulable":true}}`),
			optsv1.PatchOptions{FieldManager: Manager},
		)
		if err != nil {
			return fmt.Errorf("failed to update node %s: %v", ns[i].name, err)
		}

		// evict all pods on the node
		for _, podName := range ns[i].pods {
			err = evictClient.Evict(
				ctx,
				&policy.Eviction{
					ObjectMeta: optsv1.ObjectMeta{
						Name:      podName,
						Namespace: cm.namespace,
					},
				},
			)
			if err != nil {
				return fmt.Errorf("failed to evict pod %s: %v", podName, err)
			}
		}
	}

	return nil
}

func (cm *simClusterManager) GetCurrentNode() (int, error) {
	ctx := context.Background()
	opts := optsv1.ListOptions{
		FieldSelector: "spec.unschedulable=false",
		LabelSelector: "type=vm",
	}
	nodes, err := cm.c.CoreV1().Nodes().List(ctx, opts)
	if err != nil {
		return 0, fmt.Errorf("failed to list nodes: %v", err)
	}
	return len(nodes.Items), nil
}

func (cm *simClusterManager) GetVmSpec(n int) spec.Spec {
	s := make(spec.Spec)
	s[spec.Cpu] = cm.vmCpu * float64(n)
	s[spec.Mem] = cm.vmMem * float64(n)
	return s
}

func (cm *simClusterManager) GetServerlessPodSpec(pods []*corev1.Pod) spec.Spec {
	n := len(pods)
	return cm.GetServerlessSpecByNum(n)
}

func (cm *simClusterManager) GetServerlessSpecByNum(n int) spec.Spec {
	s := make(spec.Spec)
	s[spec.Cpu] = cm.serverlessCpu * float64(n)
	s[spec.Mem] = cm.serverlessMem * float64(n)
	return s
}

func (cm *simClusterManager) GetServerlessPrice(pods []*corev1.Pod) float64 {
	s := cm.GetServerlessPodSpec(pods)
	return cm.GetServerlessPriceBySpec(s)
}

func (cm *simClusterManager) GetServerlessPriceBySpec(s spec.Spec) float64 {
	cpu := s[spec.Cpu]
	mem := s[spec.Mem]
	return cpu*cm.serverlessCpuPrice + mem*cm.serverlessMemPrice
}

func (cm *simClusterManager) GetVmPrice(n int) float64 {
	return cm.vmPrice * float64(n)
}
