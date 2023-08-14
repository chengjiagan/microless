package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getKubeClient() (*kubernetes.Clientset, error) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

type metric struct {
	ts               time.Time
	nodeNum          int
	serverlessPodNum int
}

func collectMetrics(ctx context.Context, c *kubernetes.Clientset) <-chan metric {
	fmt.Println("start collecting metrics")
	out := make(chan metric)
	go func() {
		defer close(out)
		metrics := make([]metric, 0)
		for ctx.Err() == nil {
			time.Sleep(1 * time.Second)

			nodeNum, err := getNodeNum(c)
			check(err)
			serverlessPodNum, err := getServerlessPodNum(c)
			check(err)
			m := metric{
				ts:               time.Now(),
				nodeNum:          nodeNum,
				serverlessPodNum: serverlessPodNum,
			}
			metrics = append(metrics, m)
		}

		for _, m := range metrics {
			out <- m
		}
	}()
	return out
}

func getNodeNum(c *kubernetes.Clientset) (int, error) {
	// get nodes
	opts := metav1.ListOptions{
		LabelSelector: "type=vm",
	}
	nodes, err := c.CoreV1().Nodes().List(context.Background(), opts)
	if err != nil {
		return 0, err
	}

	// check if node is ready
	num := 0
	for _, node := range nodes.Items {
		for _, cond := range node.Status.Conditions {
			if cond.Type == "Ready" && cond.Status == "True" {
				num++
				break
			}
		}
	}

	return num, nil
}

func getServerlessPodNum(c *kubernetes.Clientset) (int, error) {
	opts := metav1.ListOptions{
		LabelSelector: "type=serverless",
	}
	pods, err := c.CoreV1().Pods("").List(context.Background(), opts)
	if err != nil {
		return 0, err
	}
	return len(pods.Items), nil
}

// getMetrics returns all metrics from the channel
func getMetrics(ch <-chan metric) []metric {
	res := make([]metric, 0)
	for m := range ch {
		res = append(res, m)
	}
	return res
}

func saveMetrics(filename string, metrics []metric) {
	// open file
	fp, err := os.Create(filename)
	check(err)
	defer fp.Close()

	// write
	_, err = fp.WriteString("ts,nodeNum,serverlessPodNum\n")
	check(err)
	for _, m := range metrics {
		_, err = fp.WriteString(fmt.Sprintf("%v,%v,%v\n", m.ts.UnixMilli(), m.nodeNum, m.serverlessPodNum))
		check(err)
	}
}
