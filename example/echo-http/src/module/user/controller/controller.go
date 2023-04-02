package controller

import (
	dto2 "echo-http/src/module/user/dto"
	"github.com/gestgo/gest/package/core/router"
	"github.com/gestgo/gest/package/extension/echofx/parser"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type IUserController interface {
	Create()
	FindOne()
	FindAll()
	Update()
	Delete()
}
type Params struct {
	fx.In
	EchoRouter *echo.Group
	Logger     *zap.SugaredLogger
}
type userController struct {
	Params
}

func NewUserRouter(params Params) Result {
	c := &userController{
		params,
	}
	return Result{Router: router.NewBaseRouter[IUserController](c)}
}

type Result struct {
	fx.Out
	Router router.IRouter `group:"echoRouters"`
}

func (b *userController) Create() {

	b.EchoRouter.POST("/users", func(c echo.Context) error {
		body := c.Get("body").(*dto2.CreateUser)
		return c.JSON(http.StatusOK, body)
	}, parser.NewBodyParser[dto2.CreateUser]("body", true).Parser)

}

func (b *userController) FindAll() {
	b.EchoRouter.GET("/users", func(c echo.Context) error {
		param := c.Get("query").(*dto2.GetListUserQuery)
		b.Logger.Info(param)
		return c.JSON(http.StatusOK, param)
	}, parser.NewQueryParser[dto2.GetListUserQuery]("query", true).Parser)

}

func (b *userController) FindOne() {

	b.EchoRouter.GET("/users/:id", func(c echo.Context) error {

		u := c.Get("param").(*dto2.GetUserById)
		return c.JSON(http.StatusOK, u)
	}, parser.NewParamsParser[dto2.GetUserById]("param", true).Parser)

}
func (b *userController) Update() {

	b.EchoRouter.PUT("/users/:id", func(c echo.Context) error {

		u := c.Get("request").(*dto2.UpdateUser)
		return c.JSON(http.StatusOK, u)
	}, parser.NewRequestParser[dto2.UpdateUser]("request", true).Parser)
}

func (b *userController) Delete() {
	b.EchoRouter.DELETE("/users/:id", func(c echo.Context) error {
		u := c.Get("request").(*dto2.DeleteUserById)
		return c.JSON(http.StatusOK, u)
	}, parser.NewRequestParser[dto2.DeleteUserById]("request", true).Parser)
}
