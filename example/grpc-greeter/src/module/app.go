package module

import (
	"grpc-greeter/src/service"

	grpcext "github.com/gestgo/gest/package/extension/grpc"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Supply(grpcext.Config{
			Port: 50051,
		}),

		fx.Provide(grpcext.AsHandler(service.NewGreeterHandler)),

		grpcext.Module(),

		fx.Invoke(func(*grpc.Server) {}),
	)
}
