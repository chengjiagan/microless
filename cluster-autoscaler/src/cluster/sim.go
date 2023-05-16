package cluster

import (
	"context"
	"fmt"
	"microless/cluster-autoscaler/spec"

	corev1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1"
	optsv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	c         *kubernetes.Clientset
	namespace string
}

// simulating adding a new node
// actually just set the node to schedulable
func (cm *simClusterManager) ScaleUpNode(n int) error {
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
		node := &nodes.Items[i]
		node.Spec.Unschedulable = false
		opts := optsv1.UpdateOptions{
			FieldManager: Manager,
		}
		_, err := cm.c.CoreV1().Nodes().Update(ctx, node, opts)
		if err != nil {
			return fmt.Errorf("failed to update node %s: %v", node.Name, err)
		}
	}
	return nil
}

// simulating removing a node
// actually just set the node to unschedulable and evict all pods on it
func (cm *simClusterManager) ScaleDownNode(n int) error {
	ctx := context.Background()
	opts := optsv1.ListOptions{
		FieldSelector: "spec.unschedulable=false",
		LabelSelector: "type=vm",
	}
	nodes, err := cm.c.CoreV1().Nodes().List(ctx, opts)
	if err != nil {
		return fmt.Errorf("failed to list nodes: %v", err)
	}

	if len(nodes.Items) < n {
		return fmt.Errorf("not enough nodes to scale down")
	}

	// TODO: sort nodes by number of pods on it
	for i := 0; i < n; i++ {
		node := &nodes.Items[i]

		// update node as unschedulaable
		node.Spec.Unschedulable = true
		updateOpts := optsv1.UpdateOptions{
			FieldManager: Manager,
		}
		_, err := cm.c.CoreV1().Nodes().Update(ctx, node, updateOpts)
		if err != nil {
			return fmt.Errorf("failed to update node %s: %v", node.Name, err)
		}

		// evict all pods on the node
		listOpts := optsv1.ListOptions{
			FieldSelector: fmt.Sprintf("spec.nodeName=%s", node.Name),
		}
		pods, err := cm.c.CoreV1().Pods(cm.namespace).List(ctx, listOpts)
		if err != nil {
			return fmt.Errorf("failed to list pods on node %s: %v", node.Name, err)
		}
		client := cm.c.PolicyV1().Evictions(cm.namespace)
		for _, pod := range pods.Items {
			err = client.Evict(ctx, &policy.Eviction{
				ObjectMeta: optsv1.ObjectMeta{
					Name:      pod.Name,
					Namespace: pod.Namespace,
				},
			})
			if err != nil {
				return fmt.Errorf("failed to evict pod %s: %v", pod.Name, err)
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
