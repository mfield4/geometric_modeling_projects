package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Ui interface {
	Input(*MouseEvent) *Command
	Update(*Command)
	Render(*sdl.Renderer)
	Rect() *sdl.Rect
	Register(map[int]Ui)
	Layer() int
	Selected(bool)
	IsSelected() bool
}
