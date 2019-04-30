package main

import (
	"github.com/mfield4/178_projects/pkg/app"
)

func main() {
	interpolation := app.NewApp()
	interpolation.SetCurrent(NewStateA())

	for interpolation.Input() {
		interpolation.Update()
		interpolation.Render()
	}
}
