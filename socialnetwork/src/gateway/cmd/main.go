package main

import (
	"context"
	"flag"
	"log"
	"microless/socialnetwork/proto/composepost"
	"microless/socialnetwork/proto/hometimeline"
	"microless/socialnetwork/proto/socialgraph"
	"microless/socialnetwork/proto/user"
	"microless/socialnetwork/proto/usertimeline"
	"microless/socialnetwork/utils"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
	// connect to ComposerPost grpc server
	conn, err := utils.NewConn(config.Service.ComposePost)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = composepost.RegisterComposePostServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to HomeTimeline grpc server
	conn, err = utils.NewConn(config.Service.HomeTimeline)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = hometimeline.RegisterHomeTimelineServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// connect to SocialGraph grpc server
	conn, err = utils.NewConn(config.Service.SocialGraph)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = socialgraph.RegisterSocialGraphServiceHandler(ctx, mux, conn)
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
	// connect to UserTimeline grpc server
	conn, err = utils.NewConn(config.Service.UserTimeline)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = usertimeline.RegisterUserTimelineServiceHandler(ctx, mux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// setup opentelemetry http handler
	handler := otelhttp.NewHandler(mux, "gateway")

	logger.Info("start server")
	err = http.ListenAndServe(config.Rest, handler)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
