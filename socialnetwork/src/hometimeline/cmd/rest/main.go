package main

import (
	"context"
	"flag"
	"log"
	"microless/socialnetwork/utils"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"

	gw "microless/socialnetwork/proto/hometimeline"
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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// setup opentelemetry
	logger.Info("connect to jaeger")
	tp, err := utils.NewTracerProvider(ctx, "HomeTimelineRest", config.Otel)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// Cleanly shutdown and flush telemetry when the application exits.
	defer func() {
		if err := utils.Shutdown(tp, ctx); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// connect to server
	conn, err := utils.NewConn(config.Service.HomeTimeline)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	err = gw.RegisterHomeTimelineServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	logger.Info("start server")
	err = http.ListenAndServe(config.Rest, mux)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
