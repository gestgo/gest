package main

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
}

func main() {
	validate := validator.New()

	// Create a universal translator and register the "en" locale
	enLocale := en.New()
	uniTrans := ut.New(enLocale, enLocale)
	trans, _ := uniTrans.GetTranslator("en")

	// Register the "en" translation for the "required" validation tag
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// Register the "en" translations for all validation tags
	en_translations.RegisterDefaultTranslations(validate, trans)

	// Validate the user struct
	user := User{}
	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// Use custom error message
			fmt.Println(err.Translate(trans))
		}
	}
}
