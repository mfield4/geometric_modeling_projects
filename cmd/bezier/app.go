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

	m1dCommands     map[int]ui.Mouse1Down
	m1uCommands     map[int]ui.Mouse1Up
	mMotionCommands map[int]ui.MouseMotion
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
		startTime:       sdl.GetTicks(),
		endTime:         sdl.GetTicks(),
		sdlWindow:       window,
		sdlRenderer:     renderer,
		Window:          ui.NewWindow(renderer),
		m1dCommands:     map[int]ui.Mouse1Down{},
		m1uCommands:     map[int]ui.Mouse1Up{},
		mMotionCommands: map[int]ui.MouseMotion{},
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
				if e.Button != 1 {
					return true
				}

				if e.State == sdl.PRESSED {
					for id, cmd := range app.leftButtonPressed(e) {
						cmds[id] = cmd
					}

					return true
				}

				if e.State == sdl.RELEASED {
					for id, cmd := range app.leftButtonReleased(e) {
						cmds[id] = cmd
					}

					return true
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

func (app *App) RegisterCol(colM map[int]ui.Ui) {
	app.Window.RegisterCol(colM)
}

func (app *App) RegisterM1d() {
	app.Window.RegisterM1d(app.m1dCommands)
}

func (app *App) RegisterM1u() {
	app.Window.RegisterM1u(app.m1uCommands)
}

func (app *App) RegisterMM() {
	app.Window.RegisterMM(app.mMotionCommands)
}

func (app *App) leftButtonPressed(event *sdl.MouseButtonEvent) map[int]func(int) {
	cmds := map[int]func(int){}

	for id, cmd := range app.m1dCommands {
		if cmd.PressActive(event.X, event.Y) {
			cmds[id] = func(id int) {
				app.m1dCommands[id].Mouse1Down(event.X, event.Y)
			}
		}
	}

	return cmds
}

func (app *App) leftButtonReleased(event *sdl.MouseButtonEvent) map[int]func(int) {
	cmds := map[int]func(int){}

	for id, cmd := range app.m1uCommands {
		if cmd.ReleaseActive(event.X, event.Y) {
			cmds[id] = func(id int) {
				app.m1uCommands[id].Mouse1Up(event.X, event.Y)
			}
		}
	}

	return cmds
}

func (app *App) mouseMotionCommand() map[int]func(int) {
	cmds := map[int]func(int){}

	for id, cmd := range app.mMotionCommands {
		if cmd.MotionActive() {
			cmds[id] = func(id int) {
				x, y, _ := sdl.GetMouseState()
				app.mMotionCommands[id].MouseMotion(x, y)
			}
		}
	}

	return cmds
}
