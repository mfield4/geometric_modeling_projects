package state

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Concrete states will consist of a collection of embedded events, then a collection of render components.
type State interface {
	Input() ([]func(), bool)
	Update([]func())
	Render(*sdl.Renderer)
}
