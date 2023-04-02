package

import "go.uber.org/fx"
{{cookiecutter.name}}

import (
	"go.uber.org/fx"
	"{{cookiecutter.base_path}}/{{cookiecutter.name}}/controller"
	"{{cookiecutter.base_path}}/{{cookiecutter.name}}/service"
)

func Module() fx.Option {
	return fx.Module("{{cookiecutter.name}}",
		fx.Provide(
			controller.New{{cookiecutter.name_camelcase}}Router,
			service.New{{cookiecutter.name_camelcase}}Service,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
