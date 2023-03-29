package module

import (
	"go-echo-base/src/module/user"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(user.Module())
}
