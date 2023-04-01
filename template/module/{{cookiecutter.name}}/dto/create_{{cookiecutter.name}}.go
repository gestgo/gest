package dto

type Create{{cookiecutter.name_camelcase}} struct {
	Name string `json:"name" validate:"required"`
}
