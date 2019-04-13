package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"path/filepath"
)

type Window struct {
	ID     int
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

	var width int32 = 1920
	var height int32 = 1080
	canvasW := (width * 2) / 3
	canvasH := (width * 2) / 3
	menuW := width / 3
	menuH := width / 3

	background := LoadTexture(renderer, filepath.Join(s, "background.png"))

	dst := &sdl.Rect{
		X: 0,
		Y: 0,
		W: width,
		H: height,
	}

	return &Window{
		ID:         GUID(),
		layer:      0,
		canvas:     NewCanvas(canvasW, canvasH),
		menu:       NewMenu(menuW, menuH),
		background: background,
		dst:        dst,
		width:      width,
		height:     height,
	}
}

func (w *Window) Input(*sdl.MouseButtonEvent) *Command {
	return nil
}

func (w *Window) Update(cmd *Command) {
	w.canvas.Update(cmd)
	w.menu.Update(cmd)
}

func (w *Window) Render(renderer *sdl.Renderer) {
	renderer.Copy(w.background, nil, w.dst)

	w.canvas.Render(renderer)
	w.menu.Render(renderer)
}

func (w *Window) Rect() *sdl.Rect {
	return w.dst
}

func (w *Window) Register(colM map[int]Ui) {
	colM[w.ID] = w

	w.canvas.Register(colM)
	w.menu.Register(colM)
}

func (w *Window) Layer() int {
	return w.layer
}
