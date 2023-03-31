package module

import (
	"github.com/gestgo/gest/config"
	"github.com/gestgo/gest/packages/core/controller"
	"github.com/gestgo/gest/packages/techniques/echofx"
	"github.com/gestgo/gest/packages/techniques/echofx/exceptions"
	"github.com/gestgo/gest/packages/techniques/logfx"
	"github.com/gestgo/gest/src/module/user"
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
				controller.InitControllers,
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
