package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-greeter/proto/hello/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterServiceClient(conn)

	names := []string{"World", "GestGo", "gRPC"}
	for _, name := range names {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		resp, err := client.SayHello(ctx, &pb.SayHelloRequest{Name: name})
		cancel()

		if err != nil {
			log.Fatalf("SayHello(%q) error: %v", name, err)
		}
		fmt.Printf("Response: %s\n", resp.GetMessage())
	}
}
