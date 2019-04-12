package main

import (
	"fmt"
	"github.com/mfield4/178_projects/pkg/ui"
	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	sdlWindow   *sdl.Window
	sdlRenderer *sdl.Renderer

	Window ui.Window
}

func NewApp() *App {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(":3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 640, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	// mRenderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	return &App{
		sdlWindow:   window,
		sdlRenderer: renderer,
	}
}

func (app *App) Input() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.MouseMotionEvent:
			// fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", e.Timestamp, e.Type, e.Which, e.X, e.Y, e.XRel, e.YRel)
		case *sdl.MouseButtonEvent:
			// fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n", e.Timestamp, e.Type, e.Which, e.X, e.Y, e.Button, e.State)
		case *sdl.MouseWheelEvent:
			// fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n", e.Timestamp, e.Type, e.Which, e.X, e.Y)
		case *sdl.KeyboardEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n", e.Timestamp, e.Type, e.Keysym.Sym, e.Keysym.Mod, e.State, e.Repeat)
			if e.State == 1 && e.Keysym.Sym == sdl.K_r {
				println("appending r")
			}
		}
	}

	return true
}

func (app *App) Update() {

}

func (app *App) Render() {
	_ = app.sdlRenderer.Clear()

	app.sdlRenderer.SetDrawColor(255, 0, 0, 255)
	app.sdlRenderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(800),
		H: int32(640),
	})

	app.Window.Render(app.sdlRenderer)

	app.sdlRenderer.Present()
}

func main() {
	app := NewApp()

	for app.Input() {
		app.Update()
		app.Render()
	}

}
