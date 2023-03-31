package exceptions

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IEchoCustomException interface {
	ErrorHandler(err error, c echo.Context)
}

type I18nCustomException struct {
	translators map[string]ut.Translator
}

func (i *I18nCustomException) ErrorHandler(err error, c echo.Context) {
	//lang := 'en'
	code := http.StatusInternalServerError
	//if he, ok := err.(*echo.HTTPError); ok {
	//	code = he.Code
	//}
	//switch v := err.(type) {
	//case validator.ValidationErrors:
	//	errBadRequest := err.(validator.ValidationErrors)
	//	c.JSON(400, error2.NewHTTPException(400, errBadRequest.Translate(validateC.Trans), err.Error()))
	//	return
	//default:
	//	fmt.Printf("I don't know about type %T!\n", v)
	//}

	c.JSON(code, struct {
	}{})
}

func NewI18nValidationException() IEchoCustomException {
	var translators map[string]ut.Translator
	return &I18nCustomException{
		translators: translators,
	}
}
