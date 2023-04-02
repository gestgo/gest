package user

import (
	"echo-http/src/module/user/controller"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("user",
		fx.Provide(
			controller.NewUserRouter,
			//service.NewUserService,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
