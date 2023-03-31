package validate

import (
	"github.com/go-playground/validator/v10"
)

type IValidator interface {
	Validate(i any) error
}
type NestGoValidator struct {
	Validator *validator.Validate
}

func (cv *NestGoValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewNestGoValidator(validator *validator.Validate) IValidator {
	return &NestGoValidator{
		Validator: validator,
	}
}
