package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	server "microless/media/movieinfo/internal/movieinfoserver"
	pb "microless/media/proto/movieinfo"
	"microless/media/utils"

	"go.uber.org/zap"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")
var addr = flag.String("addr", os.Getenv("SERVICE_ADDR"), "address for grpc server to listen")

func main() {
	// setup logger
	logger, err := zap.NewProduction()
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
	tp, err := utils.NewTracerProvider(ctx, "MovieInfo", config.Otel)
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
	rdb, err := utils.NewRedisClient(config.Redis.MovieInfo)
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
	col := mongodb.Database(config.MongoDB.Database).Collection("movie-info")

	// connection
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup grpc
	server, err := server.NewServer(logger.Sugar(), rdb, col, config)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterMovieInfoServiceServer(grpcServer, server)

	logger.Info("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
