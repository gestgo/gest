package module

import (
	"github.com/gestgo/gest/core/router"
	"github.com/gestgo/main/app-example/config"
	"github.com/gestgo/main/app/src/module/user"
	"github.com/gestgo/main/packages/techniques/echofx"
	"github.com/gestgo/main/packages/techniques/echofx/exceptions"
	"github.com/gestgo/main/packages/techniques/logfx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
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
				fx.ResultTags(`name:"httpPort"`))),
		fx.Provide(exceptions.NewI18nValidationException),
		echofx.Module(),
		user.Module(),
		logfx.Module(),

		fx.Invoke(
			fx.Annotate(
				router.InitRouter,
				fx.ParamTags(`group:"controllers"`),
			)),
		fx.Provide(SetGlobalPrefix),
		fx.Invoke(EnableSwagger),
		fx.Invoke(EnableLogRequest),
		fx.Invoke(EnableValidationRequest),
		fx.Invoke(EnableI18nErrorHandler),
		fx.Invoke(func(*echo.Echo) {}),
	)

}
