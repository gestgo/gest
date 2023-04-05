package i18nfx

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/fx"
)

type I18nService struct {
	i18n *ut.UniversalTranslator
}

func (i *I18nService) T(lang string, params ...string) (string, error) {
	trans, found := i.i18n.GetTranslator(lang)
	if !found {
		trans = i.i18n.GetFallback()
	}
	return trans.T(params)
}

type Params struct {
	I18n *ut.UniversalTranslator "name:universalTranslator"
}

func NewI18nService(params Params) any {
	return &I18nService{
		i18n: params.I18n,
	}
}
func Module() fx.Option {
	return fx.Module("i18nfx", fx.Provide(NewUniversalTranslator, NewI18nService))
}
func StringToPluralRule(s string) locales.PluralRule {
	switch s {
	case "Unknown":
		return locales.PluralRuleUnknown
	case "Zero":
		return locales.PluralRuleZero
	case "One":
		return locales.PluralRuleOne
	case "Two":
		return locales.PluralRuleTwo
	case "Few":
		return locales.PluralRuleFew
	case "Many":
		return locales.PluralRuleMany
	case "Other":
		return locales.PluralRuleOther
	default:
		return locales.PluralRuleUnknown
	}
}
