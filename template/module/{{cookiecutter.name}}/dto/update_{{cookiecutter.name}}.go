package dto

type Update{{cookiecutter.name_camelcase}} struct {
	ID   string `json:"id" validate:"required" param:"id"`
	Name string `json:"name" validate:"required"`
}
