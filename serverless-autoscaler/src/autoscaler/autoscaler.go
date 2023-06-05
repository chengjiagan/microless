package autoscaler

import (
	"context"
	"fmt"
	"microless/serverless-autoscaler/internal/utils"
	"time"

	"github.com/go-redis/redis/v8"
	appsv1 "k8s.io/api/apps/v1"
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
	// params from config
	interval  time.Duration
	namespace string
	apps      []string
	latency   time.Duration

	rdb *redis.Client
	c   *kubernetes.Clientset
}

func NewServerlessAutoscaler(config *utils.Config) (*ServerlessAutoscaler, error) {
	c, err := utils.NewKubeClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s client: %v", err)
	}

	opts, err := redis.ParseURL(config.RedisAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis addr: %v", err)
	}
	rdb := redis.NewClient(opts)

	return &ServerlessAutoscaler{
		interval:  time.Duration(config.Interval) * time.Second,
		namespace: config.Namespace,
		apps:      config.Apps,
		latency:   time.Duration(config.Latency) * time.Second,
		rdb:       rdb,
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

	// if there are unscheduled pods and serverless HPA is disabled
	if unscheduled && replicas == 0 {
		// scale up and enable serverless HPA
		klog.Infof("enable serverless HPA for %s", app)
		err := sa.enableHPA(ctx, app)
		if err != nil {
			return fmt.Errorf("failed to enable serverless HPA: %v", err)
		}
	}

	// if there are no unscheduled pods and number of serverless pods drops to 1
	if !unscheduled && replicas == 1 {
		// scale down and disable serverless HPA
		klog.Infof("disable serverless HPA for %s", app)
		err := sa.disableHPA(ctx, app)
		if err != nil {
			return fmt.Errorf("failed to disable serverless HPA: %v", err)
		}
	}

	// nothing to do
	return nil
}

func (sa *ServerlessAutoscaler) checkUnscheduled(ctx context.Context, app string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

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
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	name := app + "-serverless"
	scale, err := sa.c.AppsV1().Deployments(sa.namespace).GetScale(ctx, name, metav1.GetOptions{})
	if err != nil {
		return 0, fmt.Errorf("failed to get scale of %s: %v", name, err)
	}
	return int(scale.Spec.Replicas), nil
}

func (sa *ServerlessAutoscaler) enableHPA(ctx context.Context, app string) error {
	pCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	name := app + "-serverless"
	_, err := sa.c.AppsV1().Deployments(sa.namespace).Patch(
		pCtx,
		name,
		types.MergePatchType,
		[]byte(`{"spec":{"replicas":1}}`),
		metav1.PatchOptions{FieldManager: Manager},
	)
	if err != nil {
		return fmt.Errorf("failed to patch scale of %s: %v", name, err)
	}

	// notify serverless HPA is enabled after the first pod is ready
	go sa.notifyEnable(ctx, app)

	return nil
}

func (sa *ServerlessAutoscaler) notifyEnable(ctx context.Context, app string) {
	// watch deployment
	name := app + "-serverless"
	watcher, err := sa.c.AppsV1().Deployments(sa.namespace).Watch(
		ctx,
		metav1.ListOptions{
			FieldSelector: "metadata.name=" + name,
		},
	)
	if err != nil {
		klog.Errorf("failed to watch %s: %v", name, err)
		return
	}
	go func() {
		time.Sleep(1 * time.Minute)
		watcher.Stop()
	}()

	// wait for the first pod is ready
	for e := range watcher.ResultChan() {
		dep, ok := e.Object.(*appsv1.Deployment)
		if !ok {
			continue
		}

		if dep.Status.ReadyReplicas > 0 {
			// latency to wait for the first pod is actually ready
			time.Sleep(sa.latency)

			err = sa.rdb.Publish(ctx, app, "true").Err()
			if err != nil {
				klog.Errorf("failed to publish %s in redis: %v", app, err)
			}
			klog.Infof("notify serverless HPA enabled for %s", app)
			return
		}
	}

	// timeout
	klog.Infof("timeout to notify serverless HPA enabled for %s", app)
}

func (sa *ServerlessAutoscaler) disableHPA(ctx context.Context, app string) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

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

	err = sa.rdb.Publish(ctx, app, "false").Err()
	if err != nil {
		return fmt.Errorf("failed to set redis key: %v", err)
	}
	klog.Infof("notify serverless HPA disabled for %s", app)
	return nil
}
