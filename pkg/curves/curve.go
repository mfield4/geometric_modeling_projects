package curves

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Curve interface {
	Add(...sdl.Point)
	Draw() // Forces an update
	Curve() []sdl.Point
	Poly() []sdl.Point
}
