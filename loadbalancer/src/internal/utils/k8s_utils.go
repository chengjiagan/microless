package utils

import (
	"context"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewKubeClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

const nsPath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

func getNamespace() string {
	ns, err := ioutil.ReadFile(nsPath)
	if err != nil || len(ns) == 0 {
		return "default"
	}
	return string(ns)
}

func GetEndpointConns(c *kubernetes.Clientset, name string, port string) []*grpc.ClientConn {
	ctx := context.Background()
	vmEndpoint, err := c.CoreV1().Endpoints(getNamespace()).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Failed to get endpoint slice: %v", err)
	}

	return endpoint2Conns(vmEndpoint, port)
}

func WatchEndpointConns(c *kubernetes.Clientset, name string, port string) chan []*grpc.ClientConn {
	ctx := context.Background()

	opts := metav1.ListOptions{
		FieldSelector: "metadata.name=" + name,
	}
	watcher, err := c.CoreV1().Endpoints(getNamespace()).Watch(ctx, opts)
	if err != nil {
		log.Fatalf("Failed to watch endpoint slice: %v", err)
	}

	ch := make(chan []*grpc.ClientConn)
	go func() {
		for e := range watcher.ResultChan() {
			if e.Type == "ADDED" || e.Type == "MODIFIED" {
				endpoint := e.Object.(*corev1.Endpoints)
				ch <- endpoint2Conns(endpoint, port)
			}
			if e.Type == "DELETED" {
				ch <- nil
			}
		}
	}()
	return ch
}

func endpoint2Conns(ep *corev1.Endpoints, port string) []*grpc.ClientConn {
	if len(ep.Subsets) == 0 {
		return nil
	}

	conns := make([]*grpc.ClientConn, 0, len(ep.Subsets[0].Addresses))
	// TODO: support multiple subsets
	for _, addr := range ep.Subsets[0].Addresses {
		addr := addr.IP + ":" + port
		conn, err := NewConn(addr)
		if err != nil {
			log.Printf("Failed to connect to %s: %v", addr, err)
			continue
		}
		conns = append(conns, conn)
	}
	return conns
}
