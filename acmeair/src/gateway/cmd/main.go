package main

import (
	"context"
	"flag"
	"log"
	"microless/acmeair/proto/bookings"
	"microless/acmeair/proto/customer"
	"microless/acmeair/proto/flights"
	"microless/acmeair/utils"
	"net/http"
	"os"

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

	flag.Parse()
	config, err := utils.ParseConfig(*configPath)
	if err != nil {
		logger.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup opentelemetry
	logger.Info("connect to jaeger")
	tp, err := utils.NewTracerProvider(ctx, "RestfulGateway", config.Otel)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer utils.ShutdownJaeger(tp, ctx, logger)

	mux := runtime.NewServeMux()
	// connect to Bookings grpc server
	conn, err := utils.NewConn(config.Service.Bookings)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = bookings.RegisterBookingsServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to Customer grpc server
	conn, err = utils.NewConn(config.Service.Customer)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = customer.RegisterCustomerServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to flights grpc server
	conn, err = utils.NewConn(config.Service.Flights)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = flights.RegisterFlightsServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// start http server
	logger.Info("start http server")
	err = http.ListenAndServe(config.Rest, mux)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
