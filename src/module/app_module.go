package module

import (
	"github.com/labstack/echo/v4"
	"github.com/phongthien99/nest-go/config"
	"github.com/phongthien99/nest-go/packages/core/controller"
	"github.com/phongthien99/nest-go/packages/echofx"
	"github.com/phongthien99/nest-go/src/module/user"
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

		echofx.Module(),
		user.Module(),

		fx.Invoke(
			fx.Annotate(
				controller.InitControllers,
				fx.ParamTags(`group:"controllers"`),
			)),
		fx.Provide(SetGlobalPrefix),
		fx.Invoke(EnableSwagger),
		fx.Invoke(EnableLogRequest),
		fx.Invoke(EnableValidationRequest),
		fx.Invoke(func(*echo.Echo) {}),
	)

}
