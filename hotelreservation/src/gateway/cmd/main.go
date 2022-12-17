package main

import (
	"context"
	"flag"
	"log"
	"microless/hotelreservation/proto/profile"
	"microless/hotelreservation/proto/rate"
	"microless/hotelreservation/proto/reservation"
	"microless/hotelreservation/proto/search"
	"microless/hotelreservation/proto/user"
	"microless/hotelreservation/utils"
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
	// connect to Profile grpc server
	conn, err := utils.NewConn(config.Service.Profile)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = profile.RegisterProfileServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to Rate grpc server
	conn, err = utils.NewConn(config.Service.Rate)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = rate.RegisterRateServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to Reservation grpc server
	conn, err = utils.NewConn(config.Service.Reservation)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = reservation.RegisterReservationServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to Search grpc server
	conn, err = utils.NewConn(config.Service.Search)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = search.RegisterSearchServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to User grpc server
	conn, err = utils.NewConn(config.Service.User)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = user.RegisterUserServiceHandler(ctx, mux, conn)
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
