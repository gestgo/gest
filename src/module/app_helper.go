package module

import (
	validator10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/phongthien99/nest-go/config"
	docs "github.com/phongthien99/nest-go/docs"
	"github.com/phongthien99/nest-go/packages/common/validate"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func EnableValidationRequest(e *echo.Echo) {
	v := validator10.New()
	validate.RegisterEnglish(v)

	a := validate.NewNestGoValidator(v)
	e.Validator = a

}
func EnableLogRequest(e *echo.Group) {
	//e.Validator = validate.NewNestGoValidator()

}

func EnableSwagger(e *echo.Group) {
	docs.SwaggerInfo.BasePath = "/v3"
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
func SetGlobalPrefix(e *echo.Echo) *echo.Group {
	return e.Group(config.GetConfiguration().Http.BasePath)
}
