package main

import (
	"github.com/mfield4/178_projects/pkg/ui"
)

func main() {
	app := NewApp()
	app.RegisterCol(ui.GetCollisionMap())
	app.RegisterM1d()
	app.RegisterM1u()
	app.RegisterMM()

	const TimePerFrame = 16 // ms

	for {
		cmds, ok := app.Input()
		if !ok {
			break
		}
		cmds = app.Collisions(cmds)
		app.Update(cmds)
		app.Render()
		app.Delay(TimePerFrame)
	}

}
