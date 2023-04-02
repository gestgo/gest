package module

import (
	"echo-http/config"
	"echo-http/src/module/user"
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/technique/logfx"
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
				fx.ResultTags(`name:"platformEchoPort"`)),
			fx.Annotate(
				SetGlobalPrefix,
				fx.ParamTags(`name:"platformEcho"`),
			)),
		echofx.Module(),
		user.Module(),
		logfx.Module(),
		//fx.Invoke(EnableSwagger),
		fx.Invoke(EnableLogRequest),
		//fx.Invoke(EnableValidationRequest),
		//fx.Invoke(EnableErrorHandler),
		fx.Invoke(func(*echo.Echo) {}),
	)

}
