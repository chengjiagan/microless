module microless/acmeair/bookings

go 1.18

replace microless/acmeair/proto => ../proto

replace microless/acmeair/utils => ../utils

require (
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	go.mongodb.org/mongo-driver v1.11.1
	go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache v0.34.0
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
	microless/acmeair/proto v0.0.0-00010101000000-000000000000
	microless/acmeair/utils v0.0.0-00010101000000-000000000000
)

require (
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.33.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.33.0 // indirect
	go.opentelemetry.io/otel v1.9.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.9.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.9.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.9.0 // indirect
	go.opentelemetry.io/otel/sdk v1.9.0 // indirect
	go.opentelemetry.io/otel/trace v1.9.0 // indirect
	go.opentelemetry.io/proto/otlp v0.18.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sync v0.0.0-20220819030929-7fc1605a5dde // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
)
