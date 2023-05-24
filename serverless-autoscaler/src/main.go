package main

import (
	"flag"
	"microless/serverless-autoscaler/autoscaler"
	"microless/serverless-autoscaler/internal/utils"
	"os"

	"k8s.io/klog/v2"
)

var configPath = flag.String("config", os.Getenv("SERVERLESS_AUTOSCALER_CONFIG"), "path to config file")

func main() {
	flag.Parse()
	config, err := utils.ParseConfig(*configPath)
	if err != nil {
		klog.Fatalf("Failed to parse config: %v", err)
	}

	as, err := autoscaler.NewServerlessAutoscaler(config)
	if err != nil {
		klog.Fatalf("Failed to create autoscaler: %v", err)
	}

	klog.Info("Starting serverless-autoscaler")
	err = as.Run()
	if err != nil {
		klog.Fatalf("Failed to run autoscaler: %v", err)
	}
}
