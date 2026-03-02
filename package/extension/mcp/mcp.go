package mcp

import (
	"github.com/gestgo/gest/package/core/lifecycle"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"go.uber.org/fx"
)

// Params là Fx input cho RegisterMCPHooks
type Params struct {
	fx.In

	Config   Config
	Handlers []IMCPHandler `group:"mcpHandlers"`
}

// RegisterMCPHooks khởi tạo MCPAdapter và đăng ký với Fx lifecycle
//
// Trả về *mcpserver.MCPServer để các provider khác có thể inject nếu cần
// thêm tools/resources sau khi khởi tạo.
//
// Example (trong app module):
//
//	fx.Provide(
//	    mcp.AsHandler(NewFileTools),
//	    mcp.RegisterMCPHooks,
//	)
func RegisterMCPHooks(lc fx.Lifecycle, params Params) *mcpserver.MCPServer {
	adapter := New(params.Config, params.Handlers...)
	adapter.RegisterLifecycle(lc, adapter)
	return adapter.Server()
}

// Module trả về Fx module cho MCP server
func Module() fx.Option {
	return fx.Module("mcp", fx.Provide(RegisterMCPHooks))
}

// Ensure MCPAdapter implements lifecycle.AdapterLifecycle at compile time
var _ lifecycle.AdapterLifecycle = (*MCPAdapter)(nil)
