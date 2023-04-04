package loader

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type II18nLoader interface {
	LoadData() map[string]Translation
}

type I18nJsonLoader struct {
	path string
	data map[string]Translation
}

func (i *I18nJsonLoader) LoadData() map[string]Translation {
	return i.data
}

type Params struct {
	Path string
}

type Translation struct {
	Key      string `json:"key"`
	Trans    string `json:"trans"`
	Type     string `json:"type"`
	Rule     string `json:"rule"`
	Override string `json:"override"`
}

func (i *I18nJsonLoader) loadTranslations(path string) error {
	translations := make(map[string]Translation)

	files, err := filepath.Glob(filepath.Join(i.path, "*.json"))
	if err != nil {
		return err
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		var trans []Translation
		if err := json.Unmarshal(data, &trans); err != nil {
			return err
		}

		for _, t := range trans {
			translations[t.Key] = t
		}
	}
	i.data = translations
	return nil
}
func NewI18nJsonLoader(param Params) II18nLoader {
	return &I18nJsonLoader{
		path: param.Path,
		data: map[string]Translation{},
	}
}
