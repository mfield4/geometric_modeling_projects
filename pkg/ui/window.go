package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"path/filepath"
)

var WindowWidth int32 = 2560
var WindowHeight int32 = 1440

type Window struct {
	Id     int
	layer  int
	canvas *Canvas
	menu   *Menu

	background    *sdl.Texture
	dst           *sdl.Rect
	width, height int32
}

func NewWindow(renderer *sdl.Renderer) *Window {
	s, ok := os.LookupEnv("RESOURCES_DIR")
	if !ok {
		log.Fatalf("Need environment var RESOURCES_DIR\n")
	}

	canvasW := (WindowWidth * 2) / 3
	menuW := WindowWidth / 3

	background := LoadTexture(renderer, filepath.Join(s, "background.png"))

	dst := &sdl.Rect{
		X: 0,
		Y: 0,
		W: WindowWidth,
		H: WindowHeight,
	}

	canvas := NewCanvas(canvasW, WindowHeight)
	menu := NewMenu(menuW, WindowHeight, canvas)

	win := &Window{
		Id:         GUID(),
		layer:      0,
		canvas:     canvas,
		menu:       menu,
		background: background,
		dst:        dst,
		width:      WindowWidth,
		height:     WindowHeight,
	}

	win.RegisterCol()

	return win
}

func (w *Window) Render(renderer *sdl.Renderer) {
	renderer.Copy(w.background, nil, w.dst)

	w.canvas.Render(renderer)
	w.menu.Render(renderer)
}

func (w *Window) Rect() *sdl.Rect {
	return w.dst
}

func (w *Window) RegisterCol() {
	GetCollisionMap()[w.Id] = w
}

func (w *Window) Layer() int {
	return w.layer
}
