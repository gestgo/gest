package mcp

import (
	"context"
	"fmt"
	"os"

	"github.com/gestgo/gest/package/core/lifecycle"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// TransportType xác định kiểu transport của MCP server
type TransportType string

const (
	// TransportStdio dùng stdin/stdout — phù hợp cho Claude Desktop, CLI tools
	TransportStdio TransportType = "stdio"
	// TransportSSE dùng HTTP Server-Sent Events — phù hợp cho web clients
	TransportSSE TransportType = "sse"
)

// Config cấu hình cho MCPAdapter
type Config struct {
	// Name tên MCP server (hiển thị với LLM client)
	Name string
	// Version phiên bản server
	Version string
	// Transport kiểu transport: stdio hoặc sse (mặc định stdio)
	Transport TransportType
	// Port cổng lắng nghe — chỉ dùng khi Transport == TransportSSE
	Port int
}

// MCPAdapter implement lifecycle.AdapterLifecycle cho MCP server
// Nhúng BaseAdapter[Config] để tái sử dụng pattern lifecycle chung
type MCPAdapter struct {
	lifecycle.BaseAdapter[Config]

	server     *mcpserver.MCPServer
	sseServer  *mcpserver.SSEServer
	handlers   []IMCPHandler
	cancelStdio context.CancelFunc
}

// New khởi tạo MCPAdapter
//
// Parameters:
//   - cfg: Cấu hình server (name, version, transport, port)
//   - handlers: Danh sách IMCPHandler để đăng ký tools/resources/prompts
//
// Returns:
//   - *MCPAdapter đã khởi tạo, chưa start
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

// Server trả về MCPServer bên trong — dùng để đăng ký tools thủ công nếu cần
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
		go stdio.Listen(stdioCtx, os.Stdin, os.Stdout) //nolint:errcheck
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
