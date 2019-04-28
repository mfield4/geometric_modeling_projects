package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Ui interface {
	Render(*sdl.Renderer)
}
