package loadbalancerserver

import (
	"context"
	"microless/loadbalancer/utils"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *LoadBalancer) clearCurrentRate() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		s.currentRete.Store(0)
	}
}

func (s *LoadBalancer) updateRateLimit(config *utils.DeploymentConfig) {
	client := s.kubeClient.AppsV1().Deployments(config.Namespace)

	// setup initial rate limit
	deployment, err := client.Get(context.Background(), config.Name, metav1.GetOptions{})
	if err != nil {
		s.logger.Warnw("Failed to get deployment", "err", err)
	} else {
		rateLimit := deployment.Status.ReadyReplicas * s.ratePerPod
		s.rateLimit.Store(rateLimit)
		s.logger.Infow("Set initial rate limit", "rateLimit", rateLimit)
	}

	// update rate limit when deployment changes
	opts := metav1.ListOptions{
		FieldSelector: "metadata.name=" + config.Name,
		Watch:         true,
	}
	watcher, err := client.Watch(context.Background(), opts)
	if err != nil {
		s.logger.Errorw("Failed to watch deployment", "err", err)
		return
	}

	for e := range watcher.ResultChan() {
		var rateLimit int32
		if e.Type == "MODIFIED" || e.Type == "ADDED" {
			deployment, ok := e.Object.(*appsv1.Deployment)
			if !ok {
				s.logger.Warnw("Failed to convert to deployment", "err", err)
				continue
			}
			rateLimit = deployment.Status.ReadyReplicas * s.ratePerPod
		} else if e.Type == "DELETED" {
			rateLimit = 0
			s.logger.Debugw("Deployment deleted", "name", config.Name)
		} else {
			s.logger.Debugw("Skip event", "type", e.Type)
			continue
		}

		s.rateLimit.Store(rateLimit)
		s.logger.Infow("Update rate limit", "rateLimit", rateLimit)
	}
}
