package main

import (
	"github.com/mfield4/178_projects/pkg/app"
)

func main() {
	interpolation := app.NewApp(NewStateA())

	for interpolation.Input() {
		interpolation.Collision()
		interpolation.Update()
		interpolation.Render()
	}
}
