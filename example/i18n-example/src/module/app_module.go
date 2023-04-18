package module

import (
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/technique/logfx"
	"github.com/go-playground/locales/en"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"i18n-example/config"
	"i18n-example/src/module/i18nfx"
	"i18n-example/src/module/i18nfx/loader"
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
					return loader.NewI18nJsonLoader(loader.Params{Path: "/home/phongthien/Desktop/start-up/gest/example/i18n-example/locales/en"})

				},
				fx.ResultTags(`name:"i18nLoader"`),
			),
			fx.Annotate(
				en.New,
				fx.ResultTags(`group:"translators"`),
			),
		),

		echofx.Module(),
		user.Module(),
		logfx.Module(),
		i18nfx.Module(),
		fx.Invoke(func(*echo.Echo) {}),
	)

}
