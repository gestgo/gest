package user

import (
	"echo-http/src/module/user/service"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("user",
		fx.Provide(
			//controller.NewRouter,
			service.NewUserService,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
