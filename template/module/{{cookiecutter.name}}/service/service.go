package service

import "go.uber.org/fx"

type I{{cookiecutter.name_camelcase}}Service interface {
}

type {{cookiecutter.name_camelcase}}Service struct {
}

type {{cookiecutter.name_camelcase}}ServiceParams struct {
	fx.In
}

func New{{cookiecutter.name_camelcase}}Service(params {{cookiecutter.name_camelcase}}ServiceParams) I{{cookiecutter.name_camelcase}}Service {
	return &{{cookiecutter.name_camelcase}}Service{}

}
