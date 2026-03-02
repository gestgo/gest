package module

import (
	"mcp-filesystem/src/tools"

	"github.com/gestgo/gest/package/extension/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewApp() *fx.App {
	return fx.New(
		// Suppress Fx default logger (MCP uses stdio — must not print to stdout)
		fx.WithLogger(func() fxevent.Logger { return fxevent.NopLogger }),

		// MCP server config — stdio transport for Claude Desktop / LLM clients
		fx.Supply(mcp.Config{
			Name:      "filesystem",
			Version:   "1.0.0",
			Transport: mcp.TransportStdio,
		}),

		// Register FileSystemTools into group "mcpHandlers"
		fx.Provide(mcp.AsHandler(tools.NewFileSystemTools)),

		// Load MCP module (init server + register lifecycle hooks)
		mcp.Module(),

		// Keep server alive
		fx.Invoke(func(*mcpserver.MCPServer) {}),
	)
}
