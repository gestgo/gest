package main

import (
	"{{cookiecutter.project_name}}/src/module"
)

func main() {
	app := module.NewApp()
	app.Run()
}
