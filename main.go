package main

import (
	"github.com/hammi85/swerve/src/app"
)

func main() {
	application := app.NewApplication()

	application.Setup()

	application.Run()
}
