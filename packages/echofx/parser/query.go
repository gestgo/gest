package parser

import (
	"github.com/labstack/echo/v4"
)

type BindPathParams interface {
	BindPathParams(c echo.Context, i interface{}) error
}
type PathParams[T any] struct {
	name     string
	validate bool
	binder   BindPathParams
}

func (b *PathParams[T]) Parser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := new(T)
		err := b.binder.BindPathParams(c, data)
		if err != nil {
			return err
		}
		if b.validate {
			if err = c.Validate(data); err != nil {
				return err
			}

		}
		c.Set(b.name, data)

		err = next(c)

		return err
	}
}

func NewDefaultPathParamsParser[T any](name string, validate bool) IParser {
	return &BodyParser[T]{
		name:     name,
		binder:   &echo.DefaultBinder{},
		validate: validate,
	}
}
