package main

import (
	"context"
	"flag"
	"log"
	"microless/socialnetwork/utils"
	"net"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"

	pb "microless/socialnetwork/proto/user"
	server "microless/socialnetwork/user/internal/userserver"
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
	tp, err := utils.NewTracerProvider("User", config.Jaeger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// Cleanly shutdown and flush telemetry when the application exits.
	defer func() {
		if err := utils.Shutdown(tp, ctx); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// setup memcached
	logger.Info("connect to memcached")
	mc := otelmemcache.NewClientWithTracing(
		memcache.New(config.Memcached.User))

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
	col := mongodb.Database(config.MongoDB.Database).Collection("user")

	// connection
	lis, err := net.Listen("tcp", config.Grpc)
	if err != nil {
		logger.Sugar().Fatalw("failed to listen", "err", err)
	}

	// setup grpc
	server, err := server.NewServer(logger.Sugar(), col, mc, config)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterUserServiceServer(grpcServer, server)

	logger.Info("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
