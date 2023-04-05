package main

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"
	ut "github.com/go-playground/universal-translator"
	"log"
)

func main() {
	//var utrans ut.UniversalTranslator
	enc := en.New()
	utrans := ut.New(enc, enc, fr.New())
	env, _ := utrans.GetTranslator("en")
	env.AddCardinal("days-left", "There is {0} day left", locales.PluralRuleOne, false)
	env.AddCardinal("days-left", "There are {0} days left", locales.PluralRuleOther, false)

	fr, _ := utrans.FindTranslator("fr")
	fr.AddCardinal("days-left", "Il reste {0} jour", locales.PluralRuleOne, false)
	fr.AddCardinal("days-left", "Il reste {0} jours", locales.PluralRuleOther, false)

	err := utrans.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}
}
