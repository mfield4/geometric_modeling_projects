package canvas

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Basically same as main canvas. However, it instead raises and draws a raised curve.
type RaisedCanvas struct {
	mc  Canvas
	dst *sdl.Rect
}

func NewRaisedCanvas(cavnas Canvas) *RaisedCanvas {
	return &RaisedCanvas{
		mc: cavnas,
		dst: &sdl.Rect{
			X: 0,
			Y: 0,
			W: 0,
			H: 0,
		},
	}
}

func (*RaisedCanvas) Render(*sdl.Renderer) {
	panic("implement me")
}

func (*RaisedCanvas) Update() {
	panic("implement me")
}

func (*RaisedCanvas) Ref() Canvas {
	panic("implement me")
}
