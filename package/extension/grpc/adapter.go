package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gestgo/gest/package/core/lifecycle"
	"google.golang.org/grpc"
)

// Config holds configuration for GRPCAdapter
type Config struct {
	// Port is the listening port for the gRPC server
	Port int
	// ServerOptions are optional gRPC server options (e.g. interceptors, TLS)
	ServerOptions []grpc.ServerOption
}

// GRPCAdapter implements lifecycle.AdapterLifecycle for a gRPC server.
// It embeds BaseAdapter[Config] to reuse the common lifecycle pattern.
type GRPCAdapter struct {
	lifecycle.BaseAdapter[Config]

	server   *grpc.Server
	handlers []IGRPCHandler
	listener net.Listener
}

// New creates and returns a new GRPCAdapter.
//
// Parameters:
//   - cfg: Server configuration (port, server options)
//   - handlers: List of IGRPCHandler to register services
//
// Returns:
//   - *GRPCAdapter initialized but not yet started
func New(cfg Config, handlers ...IGRPCHandler) *GRPCAdapter {
	srv := grpc.NewServer(cfg.ServerOptions...)

	return &GRPCAdapter{
		BaseAdapter: lifecycle.BaseAdapter[Config]{Config: cfg},
		server:      srv,
		handlers:    handlers,
	}
}

// Server returns the underlying grpc.Server — useful for manually registering services if needed.
func (a *GRPCAdapter) Server() *grpc.Server {
	return a.server
}

// OnStart implements lifecycle.AdapterLifecycle.
// Registers all handlers then starts listening.
func (a *GRPCAdapter) OnStart(ctx context.Context) error {
	for _, h := range a.handlers {
		h.Register(a.server)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Config.Port))
	if err != nil {
		return fmt.Errorf("grpc: failed to listen on port %d: %w", a.Config.Port, err)
	}
	a.listener = lis

	go func() {
		log.Printf("grpc: starting gRPC server on %s", lis.Addr().String())
		if err := a.server.Serve(lis); err != nil {
			log.Printf("grpc: gRPC server stopped: %v", err)
		}
	}()

	return nil
}

// OnStop implements lifecycle.AdapterLifecycle.
// Gracefully stops the gRPC server.
func (a *GRPCAdapter) OnStop(ctx context.Context) error {
	a.server.GracefulStop()
	return nil
}
