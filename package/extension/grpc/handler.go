package grpc

import (
	"fmt"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// IGRPCHandler is the interface to register gRPC services into the server.
// Implement this interface for each service group (e.g. UserService, OrderService...).
//
// Example:
//
//	type UserServiceHandler struct{ ... }
//
//	func (h *UserServiceHandler) Register(s *grpc.Server) {
//	    pb.RegisterUserServiceServer(s, h)
//	}
type IGRPCHandler interface {
	Register(s *grpc.Server)
}

// AsHandler annotates a constructor into the Fx group "grpcHandlers".
//
// Example:
//
//	fx.Provide(
//	    grpcfx.AsHandler(NewUserServiceHandler),
//	    grpcfx.AsHandler(NewOrderServiceHandler),
//	)
func AsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(IGRPCHandler)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, "grpcHandlers")),
	)
}
