package controller

import (
	"github.com/gestgo/gest/package/core/router"
	"github.com/gestgo/gest/package/extension/echofx/parser"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"{{cookiecutter.base_path}}/{{cookiecutter.name}}/dto"
)

type I{{cookiecutter.name_camelcase}}Controller interface {
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

func NewController(params Params) I{{cookiecutter.name_camelcase}}Controller {
	return &Controller{
		router: params.Router,
		logger: params.Logger,
	}
}

func New{{cookiecutter.name_camelcase}}Router(params Params) Result {
	c := NewController(params)
	return Result{Controller: router.NewBaseRouter[I{{cookiecutter.name_camelcase}}Controller, I{{cookiecutter.name_camelcase}}Controller](c)}

}

type Result struct {
	fx.Out
	Controller router.IRouter `group:"controllers"`
}

func (b *Controller) Create() {

	b.router.POST("/{{cookiecutter.name_camelcase}}s", func(c echo.Context) error {
		body := c.Get("body").(*dto.Create{{cookiecutter.name_camelcase}})
		return c.JSON(http.StatusOK, body)
	}, parser.NewBodyParser[dto.Create{{cookiecutter.name_camelcase}}]("body", true).Parser)

}

func (b *Controller) FindAll() {
	b.router.GET("/{{cookiecutter.name_camelcase}}s", func(c echo.Context) error {
		param := c.Get("query").(*dto.GetList{{cookiecutter.name_camelcase}}Query)
		b.logger.Info(param)
		return c.JSON(http.StatusOK, param)
	}, parser.NewQueryParser[dto.GetList{{cookiecutter.name_camelcase}}Query]("query", true).Parser)

}

func (b *Controller) FindOne() {

	b.router.GET("/{{cookiecutter.name_camelcase}}s/:id", func(c echo.Context) error {

		u := c.Get("param").(*dto.Get{{cookiecutter.name_camelcase}}ById)
		return c.JSON(http.StatusOK, u)
	}, parser.NewParamsParser[dto.Get{{cookiecutter.name_camelcase}}ById]("param", true).Parser)

}
func (b *Controller) Update() {

	b.router.PUT("/{{cookiecutter.name_camelcase}}s/:id", func(c echo.Context) error {

		u := c.Get("request").(*dto.Update{{cookiecutter.name_camelcase}})
		return c.JSON(http.StatusOK, u)
	}, parser.NewRequestParser[dto.Update{{cookiecutter.name_camelcase}}]("request", true).Parser)
}

func (b *Controller) Delete() {
	b.router.DELETE("/{{cookiecutter.name_camelcase}}s/:id", func(c echo.Context) error {
		u := c.Get("request").(*dto.Delete{{cookiecutter.name_camelcase}}ById)
		return c.JSON(http.StatusOK, u)
	}, parser.NewRequestParser[dto.Delete{{cookiecutter.name_camelcase}}ById]("request", true).Parser)
}
