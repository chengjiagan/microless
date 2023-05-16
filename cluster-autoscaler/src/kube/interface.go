package kube

import (
	"microless/cluster-autoscaler/internal/utils"
	"microless/cluster-autoscaler/spec"

	corev1 "k8s.io/api/core/v1"
)

// kubernetes generic manager
type KubeManager interface {
	// get vm type pods in namespace
	GetVmPods(string) ([]*corev1.Pod, error)
	// get serverless type pods in namespace
	GetServerlessPods(string) ([]*corev1.Pod, error)
	// check if there's unscheduled pod
	HaveUnscheduled([]*corev1.Pod) bool
	// simulate pod schedule under spec, return number of unscheduled pods
	SimulateSchedule(spec.Spec, []*corev1.Pod) int
}

func NewKubeManager(config *utils.KubeConfig) (KubeManager, error) {
	c, err := utils.NewKubeClient()
	if err != nil {
		return nil, err
	}

	km := &kubeManager{
		vmSelector:         config.VmSelector,
		serverlessSelector: config.ServerlessSelector,
		c:                  c,
	}
	return km, nil
}
