package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	background    *sdl.Texture
	width, height int
}

func (window *Window) Render(renderer *sdl.Renderer) {
	renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(window.width),
		H: int32(window.height),
	})

}
