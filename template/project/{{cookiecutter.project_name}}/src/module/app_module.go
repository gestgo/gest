package module

import (
	"github.com/gestgo/gest/package/core/router"
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/extension/echofx/exceptions"
	"github.com/gestgo/gest/package/technique/logfx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"{{cookiecutter.project_name}}/config"
	"{{cookiecutter.project_name}}/src/module/user"
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
				fx.ResultTags(`name:"platformEchoPort"`))),
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
		//fx.Invoke(EnableValidationRequest),
		fx.Invoke(EnableErrorHandler),
		fx.Invoke(func(*echo.Echo) {}),
	)

}
