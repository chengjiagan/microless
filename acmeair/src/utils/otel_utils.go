package utils

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Cleanly shutdown and flush telemetry when the application exits.
func ShutdownJaeger(tp *tracesdk.TracerProvider, ctx context.Context, logger *zap.Logger) {
	// Do not make the application hang when it is shutdown.
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	// Shutdown will flush any remaining spans and shut down the exporter.
	if err := tp.Shutdown(ctx); err != nil {
		logger.Fatal(err.Error())
	}
}

// Initializes an OTLP exporter, and configures the corresponding trace and
// metric providers.
func NewTracerProvider(ctx context.Context, name, addr string) (*tracesdk.TracerProvider, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(name),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := tracesdk.NewBatchSpanProcessor(traceExporter)
	tracerProvider := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithResource(res),
		tracesdk.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider, nil
}
