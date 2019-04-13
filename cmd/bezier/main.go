package main

import (
	"github.com/mfield4/178_projects/pkg/ui"
)

func main() {
	app := NewApp()
	app.Register(ui.GetCollisionMap())

	const TimePerFrame = 16 // ms

	for {
		cmds, ok := app.Input()
		if !ok {
			break
		}
		app.Update(cmds)
		app.Render()
		app.Delay(TimePerFrame)
	}

}
