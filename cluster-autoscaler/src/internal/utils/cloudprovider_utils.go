package utils

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/alicloud"
	"k8s.io/autoscaler/cluster-autoscaler/config"
)

// mostly copied from https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/

func getDefaultResourceLimiter() *cloudprovider.ResourceLimiter {
	// build min/max maps for resources limits
	minResources := make(map[string]int64)
	maxResources := make(map[string]int64)

	// default values
	minResources[cloudprovider.ResourceNameCores] = 0
	minResources[cloudprovider.ResourceNameMemory] = 0
	maxResources[cloudprovider.ResourceNameCores] = 5000 * 64
	maxResources[cloudprovider.ResourceNameMemory] = 5000 * 64 * 20 * 1024 * 1024 * 1024

	return cloudprovider.NewResourceLimiter(minResources, maxResources)
}

func getDefaultNodeGroupDiscoveryOptions(config *Config) cloudprovider.NodeGroupDiscoveryOptions {
	// TODO
	return cloudprovider.NodeGroupDiscoveryOptions{
		NodeGroupSpecs: []string{},
	}
}

func BuildAlicloudCloudProvider(cfg *Config) cloudprovider.CloudProvider {
	rl := getDefaultResourceLimiter()
	do := getDefaultNodeGroupDiscoveryOptions(cfg)
	return alicloud.BuildAlicloud(config.AutoscalingOptions{}, do, rl)
}
