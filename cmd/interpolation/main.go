package main

import (
	"github.com/mfield4/178_projects/pkg/app"
)

func main() {
	interpolation := app.NewApp(GetStateA())

	for interpolation.Input() {
		interpolation.Update()
		interpolation.Render()
	}
}
