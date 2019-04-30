package canvas

import (
	"github.com/mfield4/178_projects/pkg/curves"
	"github.com/mfield4/178_projects/pkg/ui"
	"github.com/veandco/go-sdl2/sdl"
)

type Canvas interface {
	ui.Ui
	Update()
	Curve() curves.Curve
	Rect() sdl.Rect
}
