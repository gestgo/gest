package user

import (
	"i18n-example/module/user/service"
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
