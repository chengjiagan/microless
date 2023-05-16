package kube

import (
	"context"
	"fmt"

	"microless/cluster-autoscaler/spec"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	podv1 "k8s.io/kubernetes/pkg/api/v1/pod"
)

const (
	GB = 1024 * 1024 * 1024
)

type kubeManager struct {
	// params from config file
	vmSelector         string
	serverlessSelector string

	c *kubernetes.Clientset
}

func (km *kubeManager) GetVmPods(ns string) ([]*corev1.Pod, error) {
	ctx := context.Background()
	opts := metav1.ListOptions{
		LabelSelector: km.vmSelector,
	}
	list, err := km.c.CoreV1().Pods(ns).List(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list vm pods in namespace %s: %v", ns, err)
	}

	pods := make([]*corev1.Pod, len(list.Items))
	for i := range list.Items {
		pods[i] = &list.Items[i]
	}
	return pods, nil
}

func (km *kubeManager) GetServerlessPods(ns string) ([]*corev1.Pod, error) {
	ctx := context.Background()
	opts := metav1.ListOptions{
		LabelSelector: km.serverlessSelector,
	}
	list, err := km.c.CoreV1().Pods(ns).List(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list serverless pods in namespace %s: %v", ns, err)
	}

	pods := make([]*corev1.Pod, len(list.Items))
	for i := range list.Items {
		pods[i] = &list.Items[i]
	}
	return pods, nil
}

func (km *kubeManager) HaveUnscheduled(pods []*corev1.Pod) bool {
	for _, p := range pods {
		_, cond := podv1.GetPodCondition(&p.Status, corev1.PodScheduled)
		if cond != nil && cond.Status == corev1.ConditionFalse && cond.Reason == corev1.PodReasonUnschedulable {
			return true
		}
	}
	return false
}

// simulate scheduling pods on nodes with spec s
// return the number of pods that can be scheduled
func (km *kubeManager) SimulateSchedule(s spec.Spec, pods []*corev1.Pod) int {
	cur := make(spec.Spec)
	cur[spec.Cpu] = 0.0
	cur[spec.Mem] = 0.0

	n := 0
	for _, p := range pods {
		cpu, mem := podSpec(p)
		cur[spec.Cpu] += cpu
		cur[spec.Mem] += mem

		if cur.LessThan(s) {
			n++
		} else {
			break
		}
	}

	return n
}

// get cpu and mem requests of a pod
func podSpec(p *corev1.Pod) (float64, float64) {
	cpu := 0.0
	mem := 0.0

	for _, c := range p.Spec.Containers {
		cpu += c.Resources.Requests.Cpu().AsApproximateFloat64()
		mem += c.Resources.Requests.Memory().AsApproximateFloat64() / GB
	}

	return cpu, mem
}
