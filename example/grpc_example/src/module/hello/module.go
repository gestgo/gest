package hello

import (
	"go.uber.org/fx"
	"grpc_example/src/module/hello/controller"
	"grpc_example/src/module/hello/service"
)

func Module() fx.Option {
	return fx.Module("hello",
		fx.Provide(
			controller.NewHelloGRPController,
			service.NewHelloService,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
