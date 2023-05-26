package main

import (
	"context"
	"flag"
	"log"
	"microless/media/proto/composereview"
	"microless/media/proto/page"
	"microless/media/proto/rating"
	"microless/media/proto/user"
	"microless/media/proto/userreview"
	"microless/media/utils"
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
	// Cleanly shutdown and flush telemetry when the application exits.
	defer func() {
		if err := utils.Shutdown(tp, ctx); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	mux := runtime.NewServeMux()
	// connect to ComposeReview grpc server
	conn, err := utils.NewConn(config.Service.ComposeReview)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = composereview.RegisterComposeReviewHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to Page grpc server
	conn, err = utils.NewConn(config.Service.Page)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = page.RegisterPageServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to Rating grpc server
	conn, err = utils.NewConn(config.Service.Rating)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = rating.RegisterRatingServiceHandler(ctx, mux, conn)
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
	// connect to UserReview grpc server
	conn, err = utils.NewConn(config.Service.UserReview)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = userreview.RegisterUserReviewServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("start server")
	err = http.ListenAndServe(config.Grpc, mux)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
