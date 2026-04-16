package service

import (
	"context"
	"fmt"

	pb "grpc-greeter/proto/hello/v1"

	grpcext "github.com/gestgo/gest/package/extension/grpc"
	"google.golang.org/grpc"
)

// GreeterHandler implements IGRPCHandler and GreeterServiceServer
type GreeterHandler struct {
	pb.UnimplementedGreeterServiceServer
}

var _ grpcext.IGRPCHandler = (*GreeterHandler)(nil)

func NewGreeterHandler() *GreeterHandler {
	return &GreeterHandler{}
}

// Register implements grpcext.IGRPCHandler
func (h *GreeterHandler) Register(s *grpc.Server) {
	pb.RegisterGreeterServiceServer(s, h)
}

// SayHello implements GreeterServiceServer
func (h *GreeterHandler) SayHello(_ context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}
