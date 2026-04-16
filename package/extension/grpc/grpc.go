package grpc

import (
	"github.com/gestgo/gest/package/core/lifecycle"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// Params là Fx input cho RegisterGRPCHooks
type Params struct {
	fx.In

	Config   Config
	Handlers []IGRPCHandler `group:"grpcHandlers"`
}

// RegisterGRPCHooks khởi tạo GRPCAdapter và đăng ký với Fx lifecycle.
//
// Trả về *grpc.Server để các provider khác có thể inject nếu cần
// đăng ký thêm services sau khi khởi tạo.
//
// Example (trong app module):
//
//	fx.Provide(
//	    grpcfx.AsHandler(NewUserServiceHandler),
//	    grpcfx.RegisterGRPCHooks,
//	)
func RegisterGRPCHooks(lc fx.Lifecycle, params Params) *grpc.Server {
	adapter := New(params.Config, params.Handlers...)
	adapter.RegisterLifecycle(lc, adapter)
	return adapter.Server()
}

// Module trả về Fx module cho gRPC server
func Module() fx.Option {
	return fx.Module("grpc", fx.Provide(RegisterGRPCHooks))
}

// Ensure GRPCAdapter implements lifecycle.AdapterLifecycle at compile time
var _ lifecycle.AdapterLifecycle = (*GRPCAdapter)(nil)
