package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	server "microless/hotelreservation/geo/internal/geoserver"
	pb "microless/hotelreservation/proto/geo"
	"microless/hotelreservation/utils"

	"go.uber.org/zap"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")

const (
	appName = "Geo" // name of the application shown in jaeger
	colName = "geo" // name of mongodb collection
)

func main() {
	// setup logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	// parse config
	flag.Parse()
	config, err := utils.ParseConfig(*configPath)
	if err != nil {
		logger.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup opentelemetry
	logger.Info("connect to jaeger")
	tp, err := utils.NewTracerProvider(ctx, appName, config.Otel)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer utils.ShutdownJaeger(tp, ctx, logger)

	// setup mongodb
	logger.Info("connect to mongodb")
	mongodb, err := utils.NewMongodbClient(ctx, config.MongoDB.Url)
	if err != nil {
		logger.Fatal(err.Error())
	}
	col := mongodb.Database(config.MongoDB.Database).Collection(colName)
	defer utils.ShutdownMongodb(mongodb, ctx, logger)
	utils.CreateGeoIndex(ctx, col, "location")

	// setup redis
	logger.Info("connect to redis")
	rdb, err := utils.NewRedisClient(config.Redis.Geo)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer utils.ShutdownRedis(rdb, logger)

	// connection
	lis, err := net.Listen("tcp", config.Grpc)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup grpc
	server, err := server.NewServer(logger.Sugar(), col, rdb)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterGeoServiceServer(grpcServer, server)

	logger.Info("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
