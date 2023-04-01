package user

import (
	"github.com/gestgo/main/app/src/module/user/controller"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("user",
		fx.Provide(
			controller.NewRouter,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
