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

	return &Window{
		Id:         GUID(),
		layer:      0,
		canvas:     NewCanvas(canvasW, WindowHeight),
		menu:       NewMenu(menuW, WindowHeight),
		background: background,
		dst:        dst,
		width:      WindowWidth,
		height:     WindowHeight,
	}
}

func (w *Window) Render(renderer *sdl.Renderer) {
	renderer.Copy(w.background, nil, w.dst)

	w.canvas.Render(renderer)
	w.menu.Render(renderer)
}

func (w *Window) Rect() *sdl.Rect {
	return w.dst
}

func (w *Window) RegisterCol(colM map[int]Ui) {
	colM[w.Id] = w

	w.canvas.RegisterCol(colM)
	w.menu.Register(colM)
}

func (w *Window) RegisterM1d(m1d map[int]Mouse1Down) {
	w.canvas.RegisterM1d(m1d)
	w.menu.RegisterM1d(m1d)
}

func (w *Window) RegisterM1u(up map[int]Mouse1Up) {
	w.canvas.RegisterM1u(up)
}

func (w *Window) RegisterMM(colM map[int]MouseMotion) {
	w.canvas.RegisterMM(colM)
}

func (w *Window) Layer() int {
	return w.layer
}
