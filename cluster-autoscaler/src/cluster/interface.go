package cluster

import (
	"microless/cluster-autoscaler/internal/utils"
	"microless/cluster-autoscaler/spec"
	"time"

	corev1 "k8s.io/api/core/v1"
)

// vender specific cluster manager
type ClusterManager interface {
	ScaleUpNode(int) error
	ScaleDownNode(int) error

	GetCurrentNode() (int, error)
	GetVmSpec(int) spec.Spec
	GetServerlessPodSpec([]*corev1.Pod) spec.Spec
	GetServerlessSpecByNum(int) spec.Spec

	// price per second
	GetServerlessPrice([]*corev1.Pod) float64
	GetServerlessPriceBySpec(spec.Spec) float64
	GetVmPrice(n int) float64
}

func NewClusterManager(config *utils.ClusterConfig) (ClusterManager, error) {
	c, err := utils.NewKubeClient()
	if err != nil {
		return nil, err
	}

	cm := &simClusterManager{
		serverlessCpu:      config.ServerlessCpu,
		serverlessMem:      config.ServerlessMem,
		vmCpu:              config.VmCpu,
		vmMem:              config.VmMem,
		serverlessCpuPrice: config.ServerlessCpuPrice,
		serverlessMemPrice: config.ServerlessMemPrice,
		vmPrice:            config.VmPrice,
		namespace:          config.Namespace,
		scaleUpLatency:     time.Duration(config.ScaleUpLatency) * time.Minute,
		c:                  c,
	}
	return cm, nil
}
