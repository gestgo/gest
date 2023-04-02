package module

import (
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/extension/grpcfx"
	"github.com/gestgo/gest/package/technique/logfx"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"grpc_example/src/module/hello"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			fx.Annotate(
				grpc.NewServer,
				fx.ResultTags(`name:"grpcServer"`)),
			fx.Annotate(
				func() int {
					return 50051
				},
				fx.ResultTags(`name:"grpcPort"`))),
		echofx.Module(),
		grpcfx.Module(),
		logfx.Module(),
		hello.Module(),
		//fx.Invoke(
		//	fx.Annotate(
		//		router.InitRouter,
		//		fx.ParamTags(`group:"controllers"`),
		//	)),
		//fx.Provide(SetGlobalPrefix),
		//fx.Invoke(EnableSwagger),
		//fx.Invoke(EnableLogRequest),
		//fx.Invoke(EnableValidationRequest),
		//fx.Invoke(EnableErrorHandler),

		fx.Invoke(func(server *grpc.Server) {}),
	)

}
