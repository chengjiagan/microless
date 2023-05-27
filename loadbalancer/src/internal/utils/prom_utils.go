package utils

import (
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func StartMetricServer(addr string, reg *prometheus.Registry) {
	mux := http.NewServeMux()
	mux.Handle(
		"/metrics",
		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
	)
	server := &http.Server{Addr: addr, Handler: mux}
	log.Printf("Start stats server at %s", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start metric server: %v", err)
	}
}

func GetServiceAndMethod(info *grpc.UnaryServerInfo) (string, string) {
	// FullMethod: /<package>.<service>/<method>
	strs := strings.Split(info.FullMethod, "/")
	fullService := strs[1]
	method := strs[2]

	// FullService: <package>.<service>
	pkgs := strings.Split(fullService, ".")
	service := pkgs[len(pkgs)-1]

	return service, method
}
