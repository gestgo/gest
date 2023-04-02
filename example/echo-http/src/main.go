package main

import (
	"echo-http/src/module"
)

func main() {
	app := module.NewApp()
	app.Run()
}
