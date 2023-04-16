package user

import (
	"go.uber.org/fx"
	"i18n-example/src/module/user/controller"
)

func Module() fx.Option {
	return fx.Module("user",
		fx.Provide(
			controller.NewRouter,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
