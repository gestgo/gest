package module

import (
	"github.com/gestgo/main/app-example/config"
	"github.com/gestgo/main/app-example/docs"
	"github.com/gestgo/main/packages/techniques/echofx/exceptions"
	"github.com/gestgo/main/packages/techniques/validate"
	validator10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func EnableValidationRequest(e *echo.Echo) {
	//
	v := validator10.New()
	//validate.RegisterEnglish(v)
	//
	customValidator := validate.NewNestGoValidator(v)
	e.Validator = customValidator

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

func EnableI18nErrorHandler(e *echo.Echo, exception exceptions.IEchoCustomException) {
	e.HTTPErrorHandler = exception.ErrorHandler
}
