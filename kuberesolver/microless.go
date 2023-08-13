package kuberesolver

import (
	"net/url"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

// current available endpoints for a target
var addressesForTarget = make(map[string]*atomic.Int32)

// return the target key for a connection
func GetTargetOfConn(cc *grpc.ClientConn) string {
	u, err := url.Parse(cc.Target())
	if err != nil {
		panic(err)
	}

	ti, err := parseResolverTarget(resolver.Target{URL: *u})
	if err != nil {
		panic(err)
	}
	if ti.serviceNamespace == "" {
		ti.serviceNamespace = getCurrentNamespaceOrDefault()
	}

	return ti.String()
}

// get the number of available endpoints for a target
func GetAddressesForTarget(key string) int {
	return int(addressesForTarget[key].Load())
}
