package module

import (
	"github.com/gestgo/gest/example/config"
	"github.com/gestgo/gest/example/docs"
	"github.com/gestgo/gest/package/extension/echofx/exceptions"
	"github.com/gestgo/gest/package/technique/validate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func EnableValidationRequest(e *echo.Echo, validator validate.IValidator) {
	e.Validator = validator

}
func EnableLogRequest(e *echo.Group) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

}

func EnableSwagger(e *echo.Group) {
	docs.SwaggerInfo.BasePath = "/v3"
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func EnableErrorHandler(e *echo.Echo, exception exceptions.IEchoCustomException) {
	e.HTTPErrorHandler = exception.ErrorHandler
}

func SetGlobalPrefix(e *echo.Echo) *echo.Group {
	return e.Group(config.GetConfiguration().Http.BasePath)
}
