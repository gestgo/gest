package main

import "grpc-greeter/src/module"

func main() {
	app := module.NewApp()
	app.Run()
}
