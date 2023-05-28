package main

import (
	"flag"
	"log"
	"microless/pingpong/utils"
	"net"
	"os"

	pb "microless/pingpong/proto/ping"

	server "microless/pingpong/ping/internal/pingserver"
)

var configPath = flag.String("config", os.Getenv("SERVICE_CONFIG"), "path to config file")
var addr = flag.String("addr", os.Getenv("SERVICE_ADDR"), "address for grpc server to listen")

func main() {
	flag.Parse()
	config, err := utils.ParseConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	// connection
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// setup grpc
	server, err := server.NewServer(config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	grpcServer := utils.NewGRPCServer()
	pb.RegisterPingServiceServer(grpcServer, server)

	log.Print("start grpc server")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
