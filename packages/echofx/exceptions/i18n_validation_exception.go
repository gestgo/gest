package exceptions

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/phongthien99/nest-go/packages/common/exceptions"
	"net/http"
)

type IEchoCustomException interface {
	ErrorHandler(err error, c echo.Context)
}

type I18nCustomException struct {
	translators map[string]ut.Translator
}

func (i *I18nCustomException) ErrorHandler(err error, c echo.Context) {
	lang := 'en'
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	//c.Logger().Error(err)
	//errorPage := fmt.Sprintf("%d.html", code)
	//nestGoError := exceptions.HTTPException[]{}
	switch v := err.(type) {
	case *exceptions.HTTPException[validator.ValidationErrors]:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	}
	c.JSON(400, struct {
	}{})
}

func NewI18nValidationException() IEchoCustomException {
	return &I18nCustomException{}
}
