package module

import (
	"github.com/gestgo/gest/package/extension/schedulefx"
	"github.com/gestgo/gest/package/technique/logfx"
	"github.com/go-co-op/gocron"
	"go.uber.org/fx"
	"i18n-example/src/module/user"
	"time"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			fx.Annotate(
				func() *gocron.Scheduler {
					return gocron.NewScheduler(time.UTC)
				},
				fx.ResultTags(`name:"platformGoCron"`)),
		),
		user.Module(),
		logfx.Module(),
		schedulefx.Module(),
		fx.Invoke(func(*gocron.Scheduler) {}),
	)

}
