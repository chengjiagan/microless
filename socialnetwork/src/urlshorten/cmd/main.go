package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	pb "microless/socialnetwork/proto/urlshorten"
	server "microless/socialnetwork/urlshorten/internal/urlshortenserver"
	"microless/socialnetwork/utils"

	"go.uber.org/zap"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")
var addr = flag.String("addr", os.Getenv("SERVICE_ADDR"), "address for grpc server to listen")

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
	tp, err := utils.NewTracerProvider(ctx, "UrlShorten", config.Otel)
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
	rdb, err := utils.NewRedisClient(config.Redis.UrlShorten)
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
	col := mongodb.Database(config.MongoDB.Database).Collection("url-shorten")
	if err := utils.CreateIndex(ctx, col, "shortened_url"); err != nil {
		logger.Fatal(err.Error())
	}

	// connection
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup grpc
	server, err := server.NewServer(logger.Sugar(), rdb, col)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterUrlShortenServiceServer(grpcServer, server)

	logger.Info("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
