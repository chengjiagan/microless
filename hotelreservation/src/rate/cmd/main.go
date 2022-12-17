package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	pb "microless/hotelreservation/proto/rate"
	server "microless/hotelreservation/rate/internal/rateserver"
	"microless/hotelreservation/utils"

	"github.com/bradfitz/gomemcache/memcache"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")

const (
	appName       = "Rate"       // name of the application shown in jaeger
	ratePlanName  = "rate-plan"  // name of mongodb collection
	hotelRateName = "hotel-rate" // name of mongodb collection
	colIndex      = "hotel_id"   // new index for mongodb collection
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
	defer utils.ShutdownMongodb(mongodb, ctx, logger)
	ratePlanCol := mongodb.Database(config.MongoDB.Database).Collection(ratePlanName)
	hotelRateCol := mongodb.Database(config.MongoDB.Database).Collection(hotelRateName)
	// create index
	err = utils.CreateIndex(ctx, hotelRateCol, colIndex, true)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup memcached
	logger.Info("connect to memcached")
	mc := otelmemcache.NewClientWithTracing(memcache.New(config.Memcached.Rate))

	// connection
	lis, err := net.Listen("tcp", config.Grpc)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup grpc
	server, err := server.NewServer(logger.Sugar(), ratePlanCol, hotelRateCol, mc)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterRateServiceServer(grpcServer, server)

	logger.Info("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
