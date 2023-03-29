package main

import "go-echo-base/src/module"

func main() {
	app := module.NewApp()
	app.Run()
}
