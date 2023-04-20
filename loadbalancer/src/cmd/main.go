package main

import (
	"flag"
	"log"
	"microless/loadbalancer/utils"
	"net"
	"net/http"
	"os"

	server "microless/loadbalancer/internal/loadbalancerserver"

	"github.com/mwitkow/grpc-proxy/proxy"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var configPath = flag.String("config", os.Getenv("LOADBALANCER_CONFIG"), "path to config file")

func main() {
	// setup logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	config, err := utils.ParseConfig(*configPath)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// connect to kubernetes api server
	kubeClient, err := utils.NewKubeClient(logger.Sugar())
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup server
	server, err := server.NewServer(logger.Sugar(), config, kubeClient)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup grpc server
	lis, err := net.Listen("tcp", config.Grpc)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := grpc.NewServer(
		grpc.UnknownServiceHandler(proxy.TransparentHandler(server.GrpcProxyHandler())),
	)
	defer grpcServer.Stop()

	// setup http server
	httpServer := &http.Server{
		Addr:    config.Http,
		Handler: server.HttpProxyHandler(),
	}
	defer httpServer.Shutdown(context.Background())

	// start server
	g := &errgroup.Group{}
	g.Go(func() error {
		logger.Info("start grpc server")
		return grpcServer.Serve(lis)
	})
	g.Go(func() error {
		logger.Info("start http server")
		return httpServer.ListenAndServe()
	})

	err = g.Wait()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
