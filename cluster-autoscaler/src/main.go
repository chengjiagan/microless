package main

import (
	"flag"
	"microless/cluster-autoscaler/autoscaler"
	"microless/cluster-autoscaler/internal/utils"
	"os"

	"k8s.io/klog/v2"
)

var configPath = flag.String("config", os.Getenv("CLUSTER_AUTOSCALER_CONFIG"), "path to config file")

func main() {
	flag.Parse()
	config, err := utils.ParseConfig(*configPath)
	if err != nil {
		klog.Fatalf("Failed to parse config: %v", err)
	}

	as, err := autoscaler.NewAutoscaler(config)
	if err != nil {
		klog.Fatalf("Failed to create autoscaler: %v", err)
	}

	klog.Info("Starting cluster-autoscaler")
	err = as.Run()
	if err != nil {
		klog.Fatalf("Failed to run autoscaler: %v", err)
	}
}
