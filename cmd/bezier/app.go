package main

import (
	"github.com/mfield4/178_projects/pkg/ui"
	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	startTime, endTime uint32

	sdlWindow   *sdl.Window
	sdlRenderer *sdl.Renderer

	Window *ui.Window
}

func NewApp() *App {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(":3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, ui.WindowWidth, ui.WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	// mRenderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	return &App{
		startTime:   sdl.GetTicks(),
		endTime:     sdl.GetTicks(),
		sdlWindow:   window,
		sdlRenderer: renderer,
		Window:      ui.NewWindow(renderer),
	}
}

func (app *App) Input() (map[int]func(int), bool) {
	cmds := map[int]func(int){}

	status := func() bool {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return false

			case *sdl.MouseButtonEvent:
				if e.Button == 3 {
					app.rightButton(e, cmds)
				}

				if e.Button == 1 {
					app.leftButton(e, cmds)
				}

				return true

			case *sdl.KeyboardEvent:
				//fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n", e.Timestamp, e.Type, e.Keysym.Sym, e.Keysym.Mod, e.State, e.Repeat)
				if e.State == 1 && e.Keysym.Sym == sdl.K_r {
					println("appending r")
				}
			}
		}
		sdl.FlushEvents(sdl.FIRSTEVENT, sdl.LASTEVENT)

		return true
	}()

	if !status {
		return nil, false
	}

	//fmt.Printf("mouse states: %d, %d\n", x, y)
	for id, cmd := range app.mouseMotionCommand() {
		if _, ok := cmds[id]; ok {
			continue
		}

		cmds[id] = cmd
	}

	return cmds, true
}

func (app *App) leftButton(e *sdl.MouseButtonEvent, cmds map[int]func(int)) {
	if e.State == sdl.PRESSED {
		for id, cmd := range app.leftButtonPressed(e) {
			cmds[id] = cmd
		}
		return
	}

	if e.State == sdl.RELEASED {
		for id, cmd := range app.leftButtonReleased(e) {
			cmds[id] = cmd
		}
		return
	}
}

func (app *App) rightButton(e *sdl.MouseButtonEvent, cmds map[int]func(int)) {
	if e.State == sdl.PRESSED {
		for id, cmd := range app.rightButtonPressed(e) {
			cmds[id] = cmd
		}
		return
	}
}

// TODO Iterates through found collisions, finding the most important one
func (app *App) Collisions(cmds map[int]func(int)) map[int]func(int) {
	layer := 0
	entities := ui.GetCollisionMap()
	for id := range cmds {
		if entities[id].Layer() > layer {
			layer = entities[id].Layer()
		}
	}

	newCommands := map[int]func(int){}

	for id, fn := range cmds {
		if entities[id].Layer() == layer {
			newCommands[id] = fn
		}
	}

	return newCommands
}

func (app *App) Update(cmds map[int]func(int)) {
	for id, fn := range cmds {
		fn(id)
	}
}

func (app *App) Render() {
	_ = app.sdlRenderer.Clear()

	app.Window.Render(app.sdlRenderer)

	app.sdlRenderer.Present()
}

func (app *App) Delay(timePerFrame uint32) {
	delta := app.endTime - app.startTime

	if delta < timePerFrame {
		sdl.Delay(timePerFrame - delta)
	}

	app.startTime = app.endTime
	app.endTime = sdl.GetTicks()
}

func (app *App) leftButtonPressed(event *sdl.MouseButtonEvent) map[int]func(int) {
	cmds := map[int]func(int){}

	for id, cmd := range ui.GetMouse1dCommands() {
		if cmd.PressActive(event.X, event.Y) {
			cmds[id] = func(id int) {
				ui.GetMouse1dCommands()[id].Mouse1Down(event.X, event.Y)
			}
		}
	}

	return cmds
}

func (app *App) leftButtonReleased(event *sdl.MouseButtonEvent) map[int]func(int) {
	cmds := map[int]func(int){}

	for id, cmd := range ui.GetMouse1uCommands() {
		if cmd.ReleaseActive(event.X, event.Y) {
			cmds[id] = func(id int) {
				ui.GetMouse1uCommands()[id].Mouse1Up(event.X, event.Y)
			}
		}
	}

	return cmds
}

func (app *App) mouseMotionCommand() map[int]func(int) {
	cmds := map[int]func(int){}

	for id, cmd := range ui.GetMouseMMCommands() {
		if cmd.MotionActive() {
			cmds[id] = func(id int) {
				x, y, _ := sdl.GetMouseState()
				ui.GetMouseMMCommands()[id].MouseMotion(x, y)
			}
		}
	}

	return cmds
}

func (app *App) rightButtonPressed(event *sdl.MouseButtonEvent) map[int]func(int) {
	cmds := map[int]func(int){}

	for id, cmd := range ui.GetMouse2dCommands() {
		if cmd.RightActive() {
			cmds[id] = func(id int) {
				ui.GetMouse2dCommands()[id].Mouse2Down(event.X, event.Y)
			}
		}
	}

	return cmds
}
