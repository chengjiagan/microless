package autoscaler

import (
	"fmt"
	"microless/cluster-autoscaler/cluster"
	"microless/cluster-autoscaler/internal/utils"
	"microless/cluster-autoscaler/kube"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
)

type Autoscaler struct {
	// params from config file
	interval       time.Duration
	stableInterval time.Duration
	namespace      string
	ratio          int

	lastScale time.Time

	km kube.KubeManager
	cm cluster.ClusterManager
}

func NewAutoscaler(config *utils.Config) (*Autoscaler, error) {
	km, err := kube.NewKubeManager(&config.Kube)
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s manager: %v", err)
	}

	cm, err := cluster.NewClusterManager(&config.Cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to create cluster manager: %v", err)
	}

	as := &Autoscaler{
		interval:       time.Duration(config.Interval) * time.Second,
		stableInterval: time.Duration(config.StableInterval) * time.Minute,
		namespace:      config.Namespace,
		ratio:          config.Ratio,
		km:             km,
		cm:             cm,
		lastScale:      time.Now(),
	}
	return as, nil
}

func (as *Autoscaler) Run() error {
	for {
		time.Sleep(as.interval)
		err := as.RunOnce()
		if err != nil {
			return err
		}
	}
}

func (as *Autoscaler) RunOnce() error {
	// do nothing if in stable interval
	if time.Since(as.lastScale) < as.stableInterval {
		klog.Info("In stable interval from last scaling")
		return nil
	}

	// get vm and serverless pods
	vm, err := as.km.GetVmPods(as.namespace)
	if err != nil {
		return fmt.Errorf("failed to get vm pods: %v", err)
	}
	serverless, err := as.km.GetServerlessPods(as.namespace)
	if err != nil {
		return fmt.Errorf("failed to get serverless pods: %v", err)
	}
	curNodes, err := as.cm.GetCurrentNode()
	if err != nil {
		return fmt.Errorf("failed to get current node: %v", err)
	}
	klog.Infof("Get %d vm pods, %d serverless pods, %d nodes", len(vm), len(serverless), curNodes)

	// if the cluster should scale up
	n := as.calcScaleUp(serverless, vm)
	if n != 0 {
		klog.Infof("Scale up %d vms", n)
		err := as.cm.ScaleUpNode(n)
		if err != nil {
			return fmt.Errorf("failed to scale up: %v", err)
		}

		// after scale up, the cluster should not scale down at once
		as.lastScale = time.Now()
		return nil
	}

	// if the cluster should scale down
	n = as.calcScaleDown(vm, curNodes)
	if n != 0 {
		klog.Infof("Scale down %d vms", n)
		err := as.cm.ScaleDownNode(n)
		if err != nil {
			return fmt.Errorf("failed to scale down: %v", err)
		}

		as.lastScale = time.Now()
		return nil
	}

	// do nothing
	klog.Info("Do nothing")
	return nil
}

// calculate the number of nodes to scale up
func (as *Autoscaler) calcScaleUp(serverlessPods []*corev1.Pod, vmPods []*corev1.Pod) int {
	// first check if vm part is overloaded
	isOverload := as.km.HaveUnscheduled(vmPods)
	if !isOverload {
		return 0
	}
	klog.Info("Vm pods are overloaded")

	// calculate serverless part's price and spec
	price := as.cm.GetServerlessPrice(serverlessPods)
	spec := as.cm.GetServerlessPodSpec(serverlessPods)

	n := 1 // number of nodes to add
	for ; n <= cluster.MaxDelta; n++ {
		// calculate the price of adding n nodes
		nprice := as.cm.GetVmPrice(n)
		// calculate the spec of adding n nodes
		nspec := as.cm.GetVmSpec(n)

		if nprice >= price || spec.LessThan(nspec) {
			break
		}
	}
	return n - 1
}

// calculate the number of nodes to scale down
func (as *Autoscaler) calcScaleDown(pods []*corev1.Pod, curNodes int) int {
	n := 1 // number of nodes to remove
	for ; n <= cluster.MaxDelta && n <= curNodes; n++ {
		// node spec of current node - n
		s := as.cm.GetVmSpec(curNodes - n)
		// number of vm pods that cannot be scheduled on vm
		nremove := len(pods) - as.km.SimulateSchedule(s, pods)
		if nremove == 0 {
			continue
		}

		nserverless := nremove * as.ratio
		serverlessSpec := as.cm.GetServerlessSpecByNum(nserverless)
		// saved cost of adding serverless pods
		serverlessPrice := as.cm.GetServerlessPriceBySpec(serverlessSpec)
		// saved cost of removing n nodes
		vmPrice := as.cm.GetVmPrice(n)
		if serverlessPrice > vmPrice {
			break
		}
	}
	return n - 1
}
