package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/phongthien99/nest-go/packages/common/exceptions"
	validateC "github.com/phongthien99/nest-go/packages/common/validate"
	"github.com/phongthien99/nest-go/packages/core/controller"
	"go.uber.org/fx"
	"net/http"
)

type ICheckController interface {
	Check()
}
type ControllerParams struct {
	fx.In
	App *echo.Echo
}
type Controller struct {
	//fx.In
	groupApi *echo.Group
}
type Result struct {
	fx.Out
	Controller controller.IController `group:"controllers""`
}

func NewController(params ControllerParams) Result {
	c := &Controller{groupApi: params.App.Group("/test")}
	return Result{Controller: controller.NewBaseController[ICheckController, *Controller](c)}

}

// FindAll godoc
// @Summary     FindAll
// @Description  Find All Alert
// @Tags         Alert
// @Param q query string false " query builder"
// @Accept       json
// @Produce      json
// @Router       /checks [get]
func (b *Controller) Check() {

	b.groupApi.POST("/abc", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(u); err != nil {
			var eb = []string{}
			//enLocale := en.New()
			//uniTrans := ut.New(enLocale, enLocale)
			//trans, _ := uniTrans.GetTranslator("en")
			for _, err := range err.(validator.ValidationErrors) {
				// Use custom error message
				fmt.Println(err.Translate(validateC.Trans))
			}
			return c.JSON(404, exceptions.NewHTTPException(400, eb, err.Error()))
			return err
		}
		return c.JSON(http.StatusOK, u)
	})

}

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,min=18,max=65"`
}
type A struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
