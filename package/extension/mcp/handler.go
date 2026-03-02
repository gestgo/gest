package mcp

import (
	"fmt"

	"go.uber.org/fx"

	mcpserver "github.com/mark3labs/mcp-go/server"
)

// IMCPHandler interface để đăng ký tools/resources/prompts vào MCP server
// Implement interface này cho mỗi nhóm tính năng (ví dụ: FileTools, DBTools...)
//
// Example:
//
//	type FileTools struct{}
//
//	func (f *FileTools) Register(s *server.MCPServer) {
//	    s.AddTool(mcp.NewTool("read_file", ...), f.handleReadFile)
//	    s.AddTool(mcp.NewTool("write_file", ...), f.handleWriteFile)
//	}
type IMCPHandler interface {
	Register(s *mcpserver.MCPServer)
}

// AsHandler annotate constructor vào Fx group "mcpHandlers"
//
// Example:
//
//	fx.Provide(
//	    mcp.AsHandler(NewFileTools),
//	    mcp.AsHandler(NewDBTools),
//	)
func AsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(IMCPHandler)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, "mcpHandlers")),
	)
}
