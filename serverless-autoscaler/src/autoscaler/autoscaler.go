package autoscaler

import (
	"context"
	"fmt"
	"microless/serverless-autoscaler/internal/utils"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	podv1 "k8s.io/kubernetes/pkg/api/v1/pod"
)

const (
	Manager = "serverless-autoscaler"
)

type ServerlessAutoscaler struct {
	interval  time.Duration
	namespace string
	apps      []string

	c *kubernetes.Clientset
}

func NewServerlessAutoscaler(config *utils.Config) (*ServerlessAutoscaler, error) {
	c, err := utils.NewKubeClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s client: %v", err)
	}

	return &ServerlessAutoscaler{
		interval:  time.Duration(config.Interval) * time.Second,
		namespace: config.Namespace,
		apps:      config.Apps,
		c:         c,
	}, nil
}

func (sa *ServerlessAutoscaler) Run() error {
	for {
		time.Sleep(sa.interval)
		err := sa.RunOnce()
		if err != nil {
			return err
		}
	}
}

func (sa *ServerlessAutoscaler) RunOnce() error {
	for _, app := range sa.apps {
		err := sa.runApp(app)
		if err != nil {
			return fmt.Errorf("failed to check app %s: %v", app, err)
		}
	}
	return nil
}

func (sa *ServerlessAutoscaler) runApp(app string) error {
	ctx := context.Background()

	// check if there are unscheduled pods
	unscheduled, err := sa.checkUnscheduled(ctx, app)
	if err != nil {
		return fmt.Errorf("failed to check unscheduled pods: %v", err)
	}
	// check if serverless HPA is replicas
	replicas, err := sa.getReplicas(ctx, app)
	if err != nil {
		return fmt.Errorf("failed to check if serverless HPA is enabled: %v", err)
	}
	klog.Infof("app %s: unscheduled=%v, replicas=%d", app, unscheduled, replicas)

	// if there are unscheduled pods and serverless HPA is disabled
	if unscheduled && replicas == 0 {
		// scale up and enable serverless HPA
		err := sa.enableHPA(ctx, app)
		if err != nil {
			return fmt.Errorf("failed to enable serverless HPA: %v", err)
		}
	}

	// if there are no unscheduled pods and number of serverless pods drops to 1
	if !unscheduled && replicas == 1 {
		// scale down and disable serverless HPA
		err := sa.disableHPA(ctx, app)
		if err != nil {
			return fmt.Errorf("failed to disable serverless HPA: %v", err)
		}
	}

	// nothing to do
	return nil
}

func (sa *ServerlessAutoscaler) checkUnscheduled(ctx context.Context, app string) (bool, error) {
	podList, err := sa.c.CoreV1().Pods(sa.namespace).List(
		ctx,
		metav1.ListOptions{
			FieldSelector: "status.phase=Pending",
			LabelSelector: "type=vm,app=" + app,
		},
	)

	if err != nil {
		return false, fmt.Errorf("failed to list pods of %s: %v", app, err)
	}

	for _, p := range podList.Items {
		_, cond := podv1.GetPodCondition(&p.Status, corev1.PodScheduled)
		if cond != nil && cond.Status == corev1.ConditionFalse && cond.Reason == corev1.PodReasonUnschedulable {
			return true, nil
		}
	}
	return false, nil
}

func (sa *ServerlessAutoscaler) getReplicas(ctx context.Context, app string) (int, error) {
	name := app + "-serverless"
	scale, err := sa.c.AppsV1().Deployments(sa.namespace).GetScale(ctx, name, metav1.GetOptions{})
	if err != nil {
		return 0, fmt.Errorf("failed to get scale of %s: %v", name, err)
	}
	return int(scale.Spec.Replicas), nil
}

func (sa *ServerlessAutoscaler) enableHPA(ctx context.Context, app string) error {
	name := app + "-serverless"
	_, err := sa.c.AppsV1().Deployments(sa.namespace).Patch(
		ctx,
		name,
		types.MergePatchType,
		[]byte(`{"spec":{"replicas":1}}`),
		metav1.PatchOptions{FieldManager: Manager},
	)
	if err != nil {
		return fmt.Errorf("failed to patch scale of %s: %v", name, err)
	}
	return nil
}

func (sa *ServerlessAutoscaler) disableHPA(ctx context.Context, app string) error {
	name := app + "-serverless"
	_, err := sa.c.AppsV1().Deployments(sa.namespace).Patch(
		ctx,
		name,
		types.MergePatchType,
		[]byte(`{"spec":{"replicas":0}}`),
		metav1.PatchOptions{FieldManager: Manager},
	)
	if err != nil {
		return fmt.Errorf("failed to patch scale of %s: %v", name, err)
	}
	return nil
}
