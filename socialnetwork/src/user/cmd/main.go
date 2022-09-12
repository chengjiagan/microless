package main

import (
	"context"
	"flag"
	"log"
	"microless/socialnetwork/utils"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	tp, err := utils.NewTracerProvider(ctx, "User", config.Otel)
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
	err = pb.RegisterUserServiceHandler(ctx, mux, conn)
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
