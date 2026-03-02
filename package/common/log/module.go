package log

import (
	"github.com/gestgo/gest/package/common/log/adapter/zap"
	"github.com/gestgo/gest/package/common/log/core"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Level string `optional:"true"`
}

func ProvideLogger(params Params) (core.ISugaredLogger, error) {
	level := core.InfoLevel
	if params.Level != "" {
		switch params.Level {
		case "DEBUG":
			level = core.DebugLevel
		case "WARN":
			level = core.WarnLevel
		case "ERROR":
			level = core.ErrorLevel
		}
	}
	return zap.NewWithConfig(zap.Config{
		Level:    level,
		Encoding: "console",
	})
}

func Module() fx.Option {
	return fx.Module("gestlog",
		fx.Provide(ProvideLogger),
	)
}
