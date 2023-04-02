package main

import (
	"grpc_example/src/module"
)

func main() {
	app := module.NewApp()
	app.Run()
}
