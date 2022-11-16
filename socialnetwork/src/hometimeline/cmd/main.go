package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	server "microless/socialnetwork/hometimeline/internal/hometimelineserver"

	"microless/socialnetwork/utils"

	pb "microless/socialnetwork/proto/hometimeline"

	"go.uber.org/zap"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")

func main() {
	// setup logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Fatal(err)
		}
	}()

	flag.Parse()
	config, err := utils.ParseConfig(*configPath)
	if err != nil {
		logger.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup opentelemetry
	logger.Info("connect to jaeger")
	tp, err := utils.NewTracerProvider(ctx, "HomeTimeline", config.Otel)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// Cleanly shutdown and flush telemetry when the application exits.
	defer func() {
		if err := utils.Shutdown(tp, ctx); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// setup redis
	logger.Info("connect to redis")
	rdb, err := utils.NewRedisClient(config.Redis.HomeTimeline)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer func() {
		if err := rdb.Close(); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// setup mongodb
	logger.Info("connect to mongodb")
	mongodb, err := utils.NewMongodbClient(ctx, config.MongoDB.Url)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer func() {
		if err := mongodb.Disconnect(context.Background()); err != nil {
			logger.Fatal(err.Error())
		}
	}()
	col := mongodb.Database(config.MongoDB.Database).Collection("home-timeline")
	if err := utils.CreateIndex(ctx, col, "user_id"); err != nil {
		logger.Fatal(err.Error())
	}

	// connection
	lis, err := net.Listen("tcp", config.Grpc)
	if err != nil {
		logger.Sugar().Fatalw("failed to listen", "err", err)
	}

	// setup grpc
	server, err := server.NewServer(logger.Sugar(), col, rdb, config)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterHomeTimelineServiceServer(grpcServer, server)

	logger.Info("start grpc server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
