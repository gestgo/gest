package mcp

import (
	"context"
	"fmt"
	"os"

	"github.com/gestgo/gest/package/core/lifecycle"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// TransportType defines the transport type of the MCP server
type TransportType string

const (
	// TransportStdio uses stdin/stdout — suitable for Claude Desktop, CLI tools
	TransportStdio TransportType = "stdio"
	// TransportSSE uses HTTP Server-Sent Events — suitable for web clients
	TransportSSE TransportType = "sse"
)

// Config holds configuration for MCPAdapter
type Config struct {
	// Name is the MCP server name (displayed to the LLM client)
	Name string
	// Version is the server version
	Version string
	// Transport is the transport type: stdio or sse (default stdio)
	Transport TransportType
	// Port is the listening port — only used when Transport == TransportSSE
	Port int
}

// MCPAdapter implements lifecycle.AdapterLifecycle for an MCP server.
// It embeds BaseAdapter[Config] to reuse the common lifecycle pattern.
type MCPAdapter struct {
	lifecycle.BaseAdapter[Config]

	server     *mcpserver.MCPServer
	sseServer  *mcpserver.SSEServer
	handlers   []IMCPHandler
	cancelStdio context.CancelFunc
}

// New creates and returns a new MCPAdapter.
//
// Parameters:
//   - cfg: Server configuration (name, version, transport, port)
//   - handlers: List of IMCPHandler to register tools/resources/prompts
//
// Returns:
//   - *MCPAdapter initialized but not yet started
func New(cfg Config, handlers ...IMCPHandler) *MCPAdapter {
	srv := mcpserver.NewMCPServer(cfg.Name, cfg.Version)

	a := &MCPAdapter{
		BaseAdapter: lifecycle.BaseAdapter[Config]{Config: cfg},
		server:      srv,
		handlers:    handlers,
	}

	if cfg.Transport == TransportSSE {
		a.sseServer = mcpserver.NewSSEServer(srv)
	}

	return a
}

// Server returns the underlying MCPServer — useful for manually registering tools if needed.
func (a *MCPAdapter) Server() *mcpserver.MCPServer {
	return a.server
}

// OnStart implement lifecycle.AdapterLifecycle
// Gọi Register trên từng handler rồi khởi động transport
func (a *MCPAdapter) OnStart(ctx context.Context) error {
	for _, h := range a.handlers {
		h.Register(a.server)
	}

	switch a.Config.Transport {
	case TransportSSE:
		go a.sseServer.Start(fmt.Sprintf(":%d", a.Config.Port)) //nolint:errcheck
	default: // stdio
		stdioCtx, cancel := context.WithCancel(context.Background())
		a.cancelStdio = cancel
		stdio := mcpserver.NewStdioServer(a.server)
		go func() {
			stdio.Listen(stdioCtx, os.Stdin, os.Stdout) //nolint:errcheck
			os.Exit(0)                                  // stdin closed → client disconnected
		}()
	}

	return nil
}

// OnStop implement lifecycle.AdapterLifecycle
// Shutdown gracefully theo transport đang dùng
func (a *MCPAdapter) OnStop(ctx context.Context) error {
	switch a.Config.Transport {
	case TransportSSE:
		if a.sseServer != nil {
			return a.sseServer.Shutdown(ctx)
		}
	default: // stdio
		if a.cancelStdio != nil {
			a.cancelStdio()
		}
	}

	return nil
}
