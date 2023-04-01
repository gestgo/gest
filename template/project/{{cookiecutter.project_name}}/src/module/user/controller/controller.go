package controller

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"{{cookiecutter.project_name}}/core/router"
	"{{cookiecutter.project_name}}/extension/echofx/parser"
	"{{cookiecutter.project_name}}/src/module/user/dto"
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
	return Result{Controller: router.NewBaseRouter[IUserController, IUserController](c)}

}

type Result struct {
	fx.Out
	Controller router.IRouter `group:"controllers"`
}

func (b *Controller) Create() {

	b.router.POST("/users", func(c echo.Context) error {
		body := c.Get("body").(*dto.CreateUser)
		return c.JSON(http.StatusOK, body)
	}, parser.NewBodyParser[dto.CreateUser]("body", true).Parser)

}

func (b *Controller) FindAll() {
	b.router.GET("/users", func(c echo.Context) error {
		param := c.Get("query").(*dto.GetListUserQuery)
		b.logger.Info(param)
		return c.JSON(http.StatusOK, param)
	}, parser.NewQueryParser[dto.GetListUserQuery]("query", true).Parser)

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
