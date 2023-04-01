package dto

type Get{{cookiecutter.name_camelcase}}ById struct {
	ID string `json:"id" validate:"required" param:"id"`
}
