package controller

import (
	"github.com/gestgo/gest/package/core/router"
	"github.com/gestgo/gest/package/extension/echofx/parser"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"i18n-example/src/module/user/dto"
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
	Router *echo.Group
	Logger *zap.SugaredLogger
}
type Controller struct {
	//fx.In
	router *echo.Group
	logger *zap.SugaredLogger
}

func NewController(params Params) IUserController {
	return &Controller{
		router: params.Router,
		logger: params.Logger,
	}
}

func NewRouter(params Params) Result {
	c := NewController(params)
	return Result{Controller: router.NewBaseRouter[IUserController](c)}

}

type Result struct {
	fx.Out
	Controller router.IRouter `group:"echoRouters"`
}

func (b *Controller) Create() {

	b.router.POST("/users", func(c echo.Context) error {
		body := c.Get("body").(*dto.CreateUser)
		return c.JSON(http.StatusOK, body)
	}, parser.NewBodyParser[dto.CreateUser]("body", true).Parser)

}

func (b *Controller) FindAll() {
	b.router.GET("/users", func(c echo.Context) error {
		//param := c.Get("query").(*dto.GetListUserQuery)
		//b.logger.Info(c.QueryParams())
		return c.String(http.StatusOK, "21321321")
	})

}

func (b *Controller) FindOne() {

	b.router.GET("/users/:id", func(c echo.Context) error {

		u := c.Get("param").(*dto.GetUserById)
		return c.JSON(http.StatusOK, u)
	}, parser.NewParamsParser[dto.GetUserById]("param", true).Parser)

}
func (b *Controller) Update() {

	b.router.PUT("/users/:id", func(c echo.Context) error {

		u := c.Get("request").(*dto.UpdateUser)
		return c.JSON(http.StatusOK, u)
	}, parser.NewRequestParser[dto.UpdateUser]("request", true).Parser)
}

func (b *Controller) Delete() {
	b.router.DELETE("/users/:id", func(c echo.Context) error {
		u := c.Get("request").(*dto.DeleteUserById)
		return c.JSON(http.StatusOK, u)
	}, parser.NewRequestParser[dto.DeleteUserById]("request", true).Parser)
}
