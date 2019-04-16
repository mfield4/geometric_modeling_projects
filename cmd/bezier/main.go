package main

func main() {
	app := NewApp()

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
