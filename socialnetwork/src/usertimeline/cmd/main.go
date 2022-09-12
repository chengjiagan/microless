package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	server "microless/socialnetwork/usertimeline/internal/usertimelineserver"

	"microless/socialnetwork/utils"

	pb "microless/socialnetwork/proto/usertimeline"

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
	tp, err := utils.NewTracerProvider(ctx, "UserTimeline", config.Otel)
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
	rdb, err := utils.NewRedisClient(config.Redis.UserTimeline)
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
	col := mongodb.Database(config.MongoDB.Database).Collection("user-timeline")
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
	pb.RegisterUserTimelineServiceServer(grpcServer, server)

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
	err = pb.RegisterUserTimelineServiceHandler(ctx, mux, conn)
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
