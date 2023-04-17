package module

import (
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/gestgo/gest/package/technique/logfx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"i18n-example/config"
	"i18n-example/src/module/user"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			fx.Annotate(
				echo.New,
				fx.ResultTags(`name:"platformEcho"`)),
			fx.Annotate(
				func() int {
					return config.GetConfiguration().Http.Port
				},
				fx.ResultTags(`name:"platformEchoPort"`)),
			fx.Annotate(
				SetGlobalPrefix,
				fx.ParamTags(`name:"platformEcho"`),
			),
			fx.Annotate(
				func() loader.II18nLoader {
					return loader.NewI18nJsonLoader(loader.Params{Path: "../../locales/en"})

				},
				fx.ParamTags(`name:"platformEcho"`),
			),
		),

		echofx.Module(),
		user.Module(),
		logfx.Module(),
		i18nfx.Module(),
		fx.Invoke(func(*echo.Echo) {}),
	)

}
