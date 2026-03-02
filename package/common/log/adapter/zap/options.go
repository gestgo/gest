package zap

import (
	"go.uber.org/zap"

	"github.com/gestgo/gest/package/common/log/core"
)

// Option is a functional option for configuring a logger
type Option func(*Config)

// WithLevel sets the log level
func WithLevel(level core.Level) Option {
	return func(c *Config) {
		c.Level = level
	}
}

// WithDevelopment enables development mode
func WithDevelopment(dev bool) Option {
	return func(c *Config) {
		c.Development = dev
	}
}

// WithEncoding sets the encoding format (json or console)
func WithEncoding(encoding string) Option {
	return func(c *Config) {
		c.Encoding = encoding
	}
}

// WithJSONEncoding sets JSON encoding
func WithJSONEncoding() Option {
	return func(c *Config) {
		c.Encoding = "json"
	}
}

// WithConsoleEncoding sets console encoding
func WithConsoleEncoding() Option {
	return func(c *Config) {
		c.Encoding = "console"
	}
}

// WithOutputPaths sets the output paths
func WithOutputPaths(paths ...string) Option {
	return func(c *Config) {
		c.OutputPaths = paths
	}
}

// WithErrorOutputPaths sets the error output paths
func WithErrorOutputPaths(paths ...string) Option {
	return func(c *Config) {
		c.ErrorOutputPaths = paths
	}
}

// NewWithOptions creates a logger with functional options
func NewWithOptions(opts ...Option) (core.ISugaredLogger, error) {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	return NewWithConfig(cfg)
}

// NewDevelopmentWithOptions creates a development logger with options
func NewDevelopmentWithOptions(opts ...Option) (core.ISugaredLogger, error) {
	cfg := DevelopmentConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	return NewWithConfig(cfg)
}

// NewProductionWithOptions creates a production logger with options
func NewProductionWithOptions(opts ...Option) (core.ISugaredLogger, error) {
	cfg := Config{
		Level:            core.InfoLevel,
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return NewWithConfig(cfg)
}

// WithZapOptions adds raw zap.Option to the logger
func WithZapOptions(_ ...zap.Option) Option {
	return func(_ *Config) {
		// Store zap options in a field (we'll need to modify Config struct)
		// For now, this is a placeholder
	}
}
