package button

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
}

func (*Button) MousePressed(button, x, y int32) {
	panic("implement me")
}

func (*Button) Render(*sdl.Renderer) {
	panic("implement me")
}

func NewButton() *Button {
	return &Button{}

}
