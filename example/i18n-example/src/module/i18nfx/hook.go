package i18nfx

import (
	"fmt"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/fx"
	"i18n-example/src/module/i18nfx/loader"
)

type I18nParams struct {
	fx.In
	Loader      loader.II18nLoader   `name:"i18nLoader"`
	Translators []locales.Translator `group:"translators"`
}

func NewUniversalTranslator(
	params I18nParams,
) Result {
	enc := en.New()
	utrans := ut.New(enc)
	AddTranslators(utrans, params.Translators)
	LoadTranslate(params, utrans)
	return Result{
		UniversalTranslator: utrans,
	}
}

type Result struct {
	fx.Out
	UniversalTranslator *ut.UniversalTranslator "name:universalTranslator"
}

func AddTranslators(utrans *ut.UniversalTranslator, translators []locales.Translator) {
	for _, translator := range translators {
		utrans.AddTranslator(translator, true)
	}

}
func LoadTranslate(params I18nParams, utrans *ut.UniversalTranslator) {
	translators := params.Translators
	data := params.Loader.LoadData()
	fmt.Println("translator", translators)
	fmt.Println("data", data)
	for _, trans := range translators {

		if val, ok := data[trans.Locale()]; ok {
			transLocale, _ := utrans.GetTranslator(trans.Locale())

			for _, translation := range val {

				switch translation.Type {
				case "Ordinal":
					transLocale.AddOrdinal(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)
					continue
				case "Cardinal":
					transLocale.AddCardinal(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)
					continue
				case "Range":
					transLocale.AddRange(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)
					continue

				default:
					transLocale.Add(translation.Key, translation.Trans, false)
					continue
				}

			}

		}

	}
}
