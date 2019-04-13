package main

import (
	"fmt"
	"github.com/mfield4/178_projects/pkg/ui"
	"github.com/veandco/go-sdl2/sdl"
	"sort"
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

	window, err := sdl.CreateWindow(":3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1920, 1080, sdl.WINDOW_SHOWN)
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

func (app *App) Input() ([]*ui.Command, bool) {
	cmds := make([]*ui.Command, 0)

	status := func() bool {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return false

			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n", e.Timestamp, e.Type, e.Which, e.X, e.Y, e.Button, e.State)
				out := app.findButtonCommand(e)
				if out == nil {
					continue
				}
				fmt.Printf("command %+v\n", out)
				cmds = append(cmds, out)
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

	x, y, _ := sdl.GetMouseState()
	//fmt.Printf("mouse states: %d, %d\n", x, y)

	cmds = append(cmds, &ui.Command{
		TypeOf: ui.MousePosition,
		Target: sdl.Point{
			X: x,
			Y: y,
		},
	})

	return cmds, true
}

func (app *App) Update(cmds []*ui.Command) {
	for _, cmd := range cmds {
		app.Window.Update(cmd)
	}
}

func (app *App) Render() {
	_ = app.sdlRenderer.Clear()

	app.Window.Render(app.sdlRenderer)

	app.sdlRenderer.Present()
}

func (app *App) Register(colM map[int]ui.Ui) {
	app.Window.Register(colM)
}

func (app *App) Delay(timePerFrame uint32) {
	delta := app.endTime - app.startTime

	if delta < timePerFrame {
		sdl.Delay(timePerFrame - delta)
	}

	app.startTime = app.endTime
	app.endTime = sdl.GetTicks()
}

func (app *App) IsSelected() bool {
	return app.Window.IsSelected()
}

func (app *App) findButtonCommand(event *sdl.MouseButtonEvent) *ui.Command {
	mousePoint := sdl.Point{
		X: event.X,
		Y: event.Y,
	}

	me := &ui.MouseEvent{
		MouseButtonEvent: event,
		MouseMotionEvent: nil,
	}

	cmds := make([]*ui.Command, 0)

	for _, ui := range ui.GetCollisionMap() {
		if r := ui.Rect(); r != nil && mousePoint.InRect(r) {
			cmd := ui.Input(me)

			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}

	return app.resolveCollisions(cmds)

}

func (app *App) findMotionCommand(event *sdl.MouseMotionEvent) *ui.Command {
	return &ui.Command{
		TypeOf: ui.MousePosition,
		Target: sdl.Point{
			X: event.X,
			Y: event.Y,
		},
		TargetId: 0,
		Layer:    0,
	}

}

// TODO Iterates through found collisions, finding the most important one
func (app *App) resolveCollisions(ints []*ui.Command) *ui.Command {
	if len(ints) == 0 {
		return nil
	}
	sort.Slice(ints, func(i, j int) bool {
		return ints[i].Layer > ints[j].Layer
	})

	return ints[0]
}
