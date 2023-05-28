package main

import (
	"context"
	"flag"
	"log"
	"microless/pingpong/proto/ping"
	"microless/pingpong/utils"

	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")
var addr = flag.String("addr", os.Getenv("SERVICE_ADDR"), "address for grpc server to listen")

func main() {
	flag.Parse()
	config, err := utils.ParseConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	// connect to Ping grpc server
	conn, err := utils.NewConn(config.Service.Ping)
	if err != nil {
		log.Fatalf("failed to connect to grpc server: %v", err)
	}
	err = ping.RegisterPingServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatalf("failed to register handler: %v", err)
	}

	log.Print("start server")
	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
