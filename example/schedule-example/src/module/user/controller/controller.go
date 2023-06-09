package controller

import (
	"github.com/gestgo/gest/package/core/router"
	"github.com/go-co-op/gocron"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type IUserController interface {
	CheckUser()
}
type Params struct {
	fx.In
	PlatformGoCron *gocron.Scheduler `name:"platformGoCron"`
	Logger         *zap.SugaredLogger
}
type Controller struct {
	//fx.In
	platformGoCron *gocron.Scheduler
	logger         *zap.SugaredLogger
}

func NewController(params Params) IUserController {
	return &Controller{
		platformGoCron: params.PlatformGoCron,
		logger:         params.Logger,
	}
}

func NewRouter(params Params) Result {
	c := NewController(params)
	return Result{Controller: router.NewBaseRouter[IUserController](c)}

}

type Result struct {
	fx.Out
	Controller router.IRouter `group:"cronJobs"`
}

func (b *Controller) CheckUser() {
	b.logger.Info("run it ")

	b.platformGoCron.Every(1).Second().Do(func() {
		b.logger.Info("test cron job")
	})

}
