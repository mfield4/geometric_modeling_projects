package main

import ap "github.com/mfield4/178_projects/pkg/app"

func main() {
	app := ap.NewApp(nil)

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
