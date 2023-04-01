package {{cookiecutter.name}}

import (
	"go.uber.org/fx"
	"{{cookiecutter.base_path}}/{{cookiecutter.name}}/controller"
	"{{cookiecutter.base_path}}/{{cookiecutter.name}}/service"
)

func Module() fx.Option {
	return fx.Module("{{cookiecutter.name}}",
		fx.Provide(
			controller.New{{cookiecutter.name}}Router,
			service.New{{cookiecutter.name}}Service,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
