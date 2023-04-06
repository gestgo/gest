package user

import (
	"go.uber.org/fx"
	"i18n-example/src/module/user/controller"
	"i18n-example/src/module/user/service"
)

func Module() fx.Option {
	return fx.Module("user",
		fx.Provide(
			controller.NewRouter,
			service.NewUserService,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
