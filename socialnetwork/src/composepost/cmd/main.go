package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	server "microless/socialnetwork/composepost/internal/composepostserver"
	pb "microless/socialnetwork/proto/composepost"
	"microless/socialnetwork/utils"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

	ctx, cancel = context.WithCancel(ctx)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// start grpc server
	go func() {
		defer cancel()
		defer wg.Done()
		logger.Info("start grpc server")
		err = utils.RunGRPCServer(ctx, grpcServer, lis)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// connect to server
	conn, err := utils.NewConn(config.Grpc)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	err = pb.RegisterComposePostServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	restServer := &http.Server{
		Addr:    config.Rest,
		Handler: mux,
	}

	// start rest server
	go func() {
		defer cancel()
		defer wg.Done()
		logger.Info("start rest server")
		err = utils.RunRestServer(ctx, restServer)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}()

	wg.Wait()
}
