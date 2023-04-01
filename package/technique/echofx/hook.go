package echofx

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	PlatformEcho *echo.Echo `name:"platformEcho"`
	HttpPort     int        `name:"platformE choPort"`
}

func RegisterEchoHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *echo.Echo {
	platformEcho := params.PlatformEcho
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				go platformEcho.Start(fmt.Sprintf(":%d", params.HttpPort))
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return platformEcho.Shutdown(ctx)

			},
		})
	return platformEcho

}
