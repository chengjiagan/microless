package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	server "microless/socialnetwork/composepost/internal/composepostserver"
	pb "microless/socialnetwork/proto/composepost"
	"microless/socialnetwork/utils"

	"go.uber.org/zap"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup opentelemetry
	logger.Info("connect to jaeger")
	tp, err := utils.NewTracerProvider(ctx, "ComposePost", config.Otel)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// Cleanly shutdown and flush telemetry when the application exits.
	defer func() {
		if err := utils.Shutdown(tp, ctx); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// connection
	lis, err := net.Listen("tcp", config.Grpc)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup grpc
	server, err := server.NewServer(logger.Sugar(), config)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterComposePostServiceServer(grpcServer, server)

	logger.Info("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
