package controller

import (
	"context"
	"github.com/gestgo/gest/package/extension/grpcfx"
	"google.golang.org/grpc"
	pb "grpc_example/src/module/hello/grpc/gen/go/hello/v1"
)

type IHelloGrpcController interface {
	grpcfx.IGrpcController
	pb.GreeterServiceServer
}
type helloGrpcController struct{}

func (s *helloGrpcController) RegisterGrpcController(server *grpc.Server) {
	pb.RegisterGreeterServiceServer(server, s)
}

func (s *helloGrpcController) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{Message: "Hello, " + req.GetName() + "!"}, nil
}

func NewHelloGRPController() grpcfx.Result {

	return grpcfx.Result{
		Controller: &helloGrpcController{},
	}
}
