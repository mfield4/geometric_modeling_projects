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

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			return nil, false
		case *sdl.MouseMotionEvent:
			// fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", e.Timestamp, e.Type, e.Which, e.X, e.Y, e.XRel, e.YRel)
		case *sdl.MouseButtonEvent:
			fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n", e.Timestamp,
				e.Type, e.Which, e.X, e.Y, e.Button, e.State)

			ui := app.findUI(e.X, e.Y)
			out := ui.Input(e)
			if out == nil {
				continue
			}
			cmds = append(cmds, out)

		case *sdl.MouseWheelEvent:
			// fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n", e.Timestamp, e.Type, e.Which, e.X, e.Y)
		case *sdl.KeyboardEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n", e.Timestamp, e.Type, e.Keysym.Sym, e.Keysym.Mod, e.State, e.Repeat)
			if e.State == 1 && e.Keysym.Sym == sdl.K_r {
				println("appending r")
			}
		}
	}

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

func (app *App) findUI(x int32, y int32) ui.Ui {
	collision := &sdl.Rect{
		X: x,
		Y: y,
		W: 1,
		H: 1,
	}

	uis := make([]ui.Ui, 0)

	for _, ui := range ui.GetCollisionMap() {
		if collision.HasIntersection(ui.Rect()) {
			uis = append(uis, ui)
		}
	}

	return app.resolveCollisions(uis)

}

// TODO Iterates through found collisions, finding the most important one
func (app *App) resolveCollisions(ints []ui.Ui) ui.Ui {
	if len(ints) == 0 {
		return nil
	}
	sort.Slice(ints, func(i, j int) bool {
		return ints[i].Layer() > ints[j].Layer()
	})

	return ints[0]
}
