package main

import (
	"github.com/gestgo/gest/example/echo-http/src/module"
)

func main() {
	app := module.NewApp()
	app.Run()
}
