package dto

type GetList{{cookiecutter.name_camelcase}}Query struct {
	Q string `json:"q" validate:"required" query:"q"`
}
