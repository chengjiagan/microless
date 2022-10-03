module microless/media/composereview

go 1.18

replace utils => ../utils

replace proto => ../proto

require (
	go.uber.org/zap v1.23.0
	golang.org/x/sync v0.0.0-20220923202941-7f9b1623fab7
	google.golang.org/protobuf v1.28.1
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
