package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// FileSystemTools implements mcp.IMCPHandler — registers all filesystem tools
type FileSystemTools struct {
	// RootDir restricts file access to this directory
	RootDir string
}

func NewFileSystemTools() *FileSystemTools {
	root, _ := os.Getwd()
	return &FileSystemTools{RootDir: root}
}

// Register implements mcp.IMCPHandler
func (f *FileSystemTools) Register(s *mcpserver.MCPServer) {
	log.Print("Registering FileSystemTools...", f.RootDir)
	s.AddTool(mcp.NewTool("read_file",
		mcp.WithDescription("Read file contents"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to the file"),
		),
	), f.readFile)

	s.AddTool(mcp.NewTool("write_file",
		mcp.WithDescription("Write content to a file (create or overwrite)"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to the file"),
		),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Content to write"),
		),
	), f.writeFile)

	s.AddTool(mcp.NewTool("list_directory",
		mcp.WithDescription("List files and directories at the given path"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Directory path"),
		),
	), f.listDirectory)

	s.AddTool(mcp.NewTool("create_directory",
		mcp.WithDescription("Create directory (including all parent directories)"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Directory path to create"),
		),
	), f.createDirectory)

	s.AddTool(mcp.NewTool("delete",
		mcp.WithDescription("Delete a file or directory"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to delete"),
		),
		mcp.WithBoolean("recursive",
			mcp.Description("Delete recursively if directory (default false)"),
		),
	), f.delete)

	s.AddTool(mcp.NewTool("move",
		mcp.WithDescription("Move or rename a file/directory"),
		mcp.WithString("source",
			mcp.Required(),
			mcp.Description("Source path"),
		),
		mcp.WithString("destination",
			mcp.Required(),
			mcp.Description("Destination path"),
		),
	), f.move)

	s.AddTool(mcp.NewTool("file_info",
		mcp.WithDescription("Get metadata of a file or directory"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to inspect"),
		),
	), f.fileInfo)
}

// ─── helpers ─────────────────────────────────────────────────────────────────

// requireString extracts a required string argument from the request
func requireString(req mcp.CallToolRequest, key string) (string, error) {
	val := mcp.ParseArgument(req, key, "")
	s, ok := val.(string)
	if !ok || s == "" {
		return "", fmt.Errorf("required parameter %q is missing", key)
	}
	return s, nil
}

// safePath resolves path and ensures it stays within RootDir
func (f *FileSystemTools) safePath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		path = filepath.Join(f.RootDir, path)
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	rootAbs, err := filepath.Abs(f.RootDir)
	if err != nil {
		return "", err
	}

	rel, err := filepath.Rel(rootAbs, abs)
	if err != nil || (len(rel) >= 2 && rel[:2] == "..") {
		return "", fmt.Errorf("path %q is outside root directory", path)
	}

	return abs, nil
}

// ─── handlers ────────────────────────────────────────────────────────────────

func (f *FileSystemTools) readFile(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := requireString(req, "path")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	safe, err := f.safePath(path)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	data, err := os.ReadFile(safe)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("read_file: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

func (f *FileSystemTools) writeFile(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := requireString(req, "path")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	content, err := requireString(req, "content")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	safe, err := f.safePath(path)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if err := os.MkdirAll(filepath.Dir(safe), 0o755); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("create parent dirs: %v", err)), nil
	}

	if err := os.WriteFile(safe, []byte(content), 0o644); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("write_file: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("written %d bytes to %s", len(content), path)), nil
}

type dirEntry struct {
	Name    string `json:"name"`
	Type    string `json:"type"` // "file" | "dir" | "symlink"
	Size    int64  `json:"size,omitempty"`
	ModTime string `json:"mod_time"`
}

func (f *FileSystemTools) listDirectory(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := requireString(req, "path")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	safe, err := f.safePath(path)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	entries, err := os.ReadDir(safe)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("list_directory: %v", err)), nil
	}

	result := make([]dirEntry, 0, len(entries))
	for _, e := range entries {
		info, _ := e.Info()
		entryType := "file"
		if e.IsDir() {
			entryType = "dir"
		} else if e.Type()&os.ModeSymlink != 0 {
			entryType = "symlink"
		}

		de := dirEntry{
			Name:    e.Name(),
			Type:    entryType,
			ModTime: info.ModTime().Format(time.RFC3339),
		}
		if !e.IsDir() {
			de.Size = info.Size()
		}
		result = append(result, de)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func (f *FileSystemTools) createDirectory(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := requireString(req, "path")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	safe, err := f.safePath(path)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if err := os.MkdirAll(safe, 0o755); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("create_directory: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("created directory: %s", path)), nil
}

func (f *FileSystemTools) delete(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := requireString(req, "path")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	recursive := mcp.ParseBoolean(req, "recursive", false)

	safe, err := f.safePath(path)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if recursive {
		err = os.RemoveAll(safe)
	} else {
		err = os.Remove(safe)
	}

	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("delete: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("deleted: %s", path)), nil
}

func (f *FileSystemTools) move(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	src, err := requireString(req, "source")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	dst, err := requireString(req, "destination")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	safeSrc, err := f.safePath(src)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	safeDst, err := f.safePath(dst)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if err := os.MkdirAll(filepath.Dir(safeDst), 0o755); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("create parent dirs: %v", err)), nil
	}

	if err := os.Rename(safeSrc, safeDst); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("move: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("moved: %s -> %s", src, dst)), nil
}

type fileInfoResult struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Size        int64  `json:"size"`
	IsDir       bool   `json:"is_dir"`
	Permissions string `json:"permissions"`
	ModTime     string `json:"mod_time"`
}

func (f *FileSystemTools) fileInfo(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, err := requireString(req, "path")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	safe, err := f.safePath(path)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	info, err := os.Stat(safe)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("file_info: %v", err)), nil
	}

	result := fileInfoResult{
		Name:        info.Name(),
		Path:        safe,
		Size:        info.Size(),
		IsDir:       info.IsDir(),
		Permissions: info.Mode().String(),
		ModTime:     info.ModTime().Format(time.RFC3339),
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}
