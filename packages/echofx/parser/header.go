package parser

import (
	"github.com/labstack/echo/v4"
)

type BindHeaders interface {
	BindHeaders(c echo.Context, i interface{}) error
}
type HeaderParser[T any] struct {
	name   string
	binder BindHeaders
}

func (b *HeaderParser[T]) Parser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := new(T)
		err := b.binder.BindHeaders(c, payload)
		if err != nil {
			return c.JSON(400, err)
		}
		c.Set(b.name, payload)

		err = next(c)

		return err
	}
}

func NewDefaultHeaderParser[T any](name string) IParser {
	return &HeaderParser[T]{
		name:   name,
		binder: &echo.DefaultBinder{},
	}
}
