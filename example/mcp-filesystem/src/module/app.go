package module

import (
	"mcp-filesystem/src/tools"

	"github.com/gestgo/gest/package/extension/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(

		// Suppress fx logs — stdio transport uses stdout for MCP protocol
		fx.NopLogger,

		// MCP server config — Stdio transport for Claude Desktop / CLI tools
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
