package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	pb "microless/hotelreservation/proto/reservation"
	server "microless/hotelreservation/reservation/internal/reservationserver"
	"microless/hotelreservation/utils"

	"github.com/bradfitz/gomemcache/memcache"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.uber.org/zap"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")

const (
	appName = "Reservation" // name of the application shown in jaeger
	colName = "reservation" // name of mongodb collection
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
	// create index
	indices := []string{"hotel_id", "in_date", "out_date"}
	for _, index := range indices {
		err = utils.CreateIndex(ctx, col, index, false)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	// setup memcached
	logger.Info("connect to memcached")
	mc := otelmemcache.NewClientWithTracing(memcache.New(config.Memcached.Reservation))

	// connection
	lis, err := net.Listen("tcp", config.Grpc)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup grpc
	server, err := server.NewServer(logger.Sugar(), col, mc, config)
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterReservationServiceServer(grpcServer, server)

	logger.Info("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
