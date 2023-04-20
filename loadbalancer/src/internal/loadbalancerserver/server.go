package loadbalancerserver

import (
	"microless/loadbalancer/utils"

	"go.uber.org/atomic"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"k8s.io/client-go/kubernetes"
)

type LoadBalancer struct {
	logger      *zap.SugaredLogger
	rateLimit   *atomic.Int32
	currentRete *atomic.Int32
	grpcConn    []*grpc.ClientConn
	httpConn    []string
	kubeClient  *kubernetes.Clientset
	ratePerPod  int32
}

func NewServer(logger *zap.SugaredLogger, config *utils.Config, kubeClient *kubernetes.Clientset) (*LoadBalancer, error) {
	vmConn, err := utils.NewConn(config.Backend.Vm)
	if err != nil {
		return nil, err
	}

	serverlessConn, err := utils.NewConn(config.Backend.Serverless)
	if err != nil {
		return nil, err
	}

	grpcConn := make([]*grpc.ClientConn, 2)
	grpcConn[VM] = vmConn
	grpcConn[SERVERLESS] = serverlessConn

	httpConn := make([]string, 2)
	httpConn[VM] = config.Backend.Vm
	httpConn[SERVERLESS] = config.Backend.Serverless

	server := &LoadBalancer{
		logger:      logger,
		grpcConn:    grpcConn,
		httpConn:    httpConn,
		rateLimit:   atomic.NewInt32(0),
		currentRete: atomic.NewInt32(0),
		kubeClient:  kubeClient,
		ratePerPod:  int32(config.RatePerPod),
	}
	go server.clearCurrentRate()
	go server.updateRateLimit(config.Deployment)

	return server, nil
}
