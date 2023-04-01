package dto

type Delete{{cookiecutter.name_camelcase}}ById struct {
	ID string `json:"id" validate:"required" param:"id"`
}
