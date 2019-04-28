package canvas

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ReducedCanvas struct {
	mc  Canvas
	dst *sdl.Rect
}

func NewReducedCanvas(canvas Canvas) *ReducedCanvas {
	return &ReducedCanvas{
		mc: canvas,
		dst: &sdl.Rect{
			X: 0,
			Y: 0,
			W: 0,
			H: 0,
		},
	}
}

func (*ReducedCanvas) Render(*sdl.Renderer) {
	panic("implement me")
}

func (*ReducedCanvas) Update() {
	panic("implement me")
}

func (*ReducedCanvas) Ref() Canvas {
	panic("implement me")
}
