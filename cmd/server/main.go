package main

import "webServerEx/internal/pkg/app"

func main() {
	application := app.NewApp()
	application.Start()
}
