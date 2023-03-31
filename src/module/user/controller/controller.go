package controller

import (
	"github.com/gestgo/gest/packages/core/controller"
	parser2 "github.com/gestgo/gest/packages/techniques/echofx/parser"
	"github.com/gestgo/gest/src/module/user/dto"
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
	return Result{Controller: controller.NewBaseController[IUserController, IUserController](c)}

}

type Result struct {
	fx.Out
	Controller controller.IController `group:"controllers""`
}

func (b *Controller) Create() {

	b.router.POST("/users", func(c echo.Context) error {
		body := c.Get("body").(*dto.CreateUser)
		return c.JSON(http.StatusOK, body)
	}, parser2.NewBodyParser[dto.CreateUser]("body", true).Parser)

}

func (b *Controller) FindAll() {
	b.router.GET("/users", func(c echo.Context) error {
		param := c.Get("query").(*dto.GetListUserQuery)
		b.logger.Info(param)
		return c.JSON(http.StatusOK, param)
	}, parser2.NewQueryParser[dto.GetListUserQuery]("query", true).Parser)

}

func (b *Controller) FindOne() {

	b.router.GET("/users/:id", func(c echo.Context) error {

		u := c.Get("param").(*dto.GetUserById)
		return c.JSON(http.StatusOK, u)
	}, parser2.NewParamsParser[dto.GetUserById]("param", true).Parser)

}
func (b *Controller) Update() {

	b.router.PUT("/users/:id", func(c echo.Context) error {

		u := c.Get("request").(*dto.UpdateUser)
		return c.JSON(http.StatusOK, u)
	}, parser2.NewRequestParser[dto.UpdateUser]("request", true).Parser)
}

func (b *Controller) Delete() {
	b.router.DELETE("/users/:id", func(c echo.Context) error {
		u := c.Get("request").(*dto.DeleteUserById)
		return c.JSON(http.StatusOK, u)
	}, parser2.NewRequestParser[dto.DeleteUserById]("request", true).Parser)
}
