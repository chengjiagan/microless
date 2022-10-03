module microless/media/rating

go 1.18

replace utils => ../utils

replace proto => ../proto

require (
	go.uber.org/zap v1.23.0
	google.golang.org/protobuf v1.28.1
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
