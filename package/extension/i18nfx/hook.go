package i18nfx

import (
	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

type I18nParams struct {
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
	for _, trans := range translators {
		if val, ok := data[trans.Locale()]; ok {
			transLocale, _ := utrans.GetTranslator(trans.Locale())
			for _, translation := range val {
				switch translation.Type {
				case "Ordinal":
					transLocale.AddOrdinal(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)
				case "Cardinal":
					transLocale.AddCardinal(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)
				case "Range":
					transLocale.AddRange(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)

				default:
					transLocale.Add(translation.Key, translation.Trans, true)
				}

			}

		}

	}
}
