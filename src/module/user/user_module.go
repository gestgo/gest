package user

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("user",
		fx.Provide(
			NewController,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
