module grpc-greeter

go 1.23.0

toolchain go1.24.2

require (
	github.com/gestgo/gest/package/extension/grpc v0.0.0-00010101000000-000000000000
	go.uber.org/fx v1.23.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

require (
	github.com/gestgo/gest/package/core/lifecycle v0.0.0-00010101000000-000000000000 // indirect
	go.uber.org/dig v1.18.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
)

replace (
	github.com/gestgo/gest/package/core/lifecycle => ../../package/core/lifecycle
	github.com/gestgo/gest/package/extension/grpc => ../../package/extension/grpc
)
