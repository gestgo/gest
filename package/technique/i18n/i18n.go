package i18n

import ut "github.com/go-playground/universal-translator"

type I18n struct {
	i18n *ut.UniversalTranslator
}

func (i *I18n) T(lang string, params ...string) (string, error) {
	trans, found := i.i18n.GetTranslator(lang)
	if !found {
		trans = i.i18n.GetFallback()
	}
	return trans.T(params)
}
