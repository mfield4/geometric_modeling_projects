package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Canvas struct {
	Id           int
	layer        int
	currentCurve int
	curves       map[int]Ui

	dst *sdl.Rect
}

func (c *Canvas) Selected(bool) {}

func (c *Canvas) IsSelected() bool {
	for _, c := range c.curves {
		if c.IsSelected() {
			return true
		}
	}

	return false
}

func NewCanvas(width int32, height int32) *Canvas {
	initCurveId := GUID()
	castel := map[int]Ui{
		initCurveId: NewCasteljauBezierCurve(initCurveId, 2),
	}

	return &Canvas{
		Id:           GUID(),
		layer:        1,
		currentCurve: initCurveId,
		curves:       castel,
		dst: &sdl.Rect{
			X: 0,
			Y: 0,
			W: width,
			H: height,
		},
	}
}

func (c *Canvas) Input(event *MouseEvent) *Command {
	// Exit conditions. Only want a left click release
	if event.Button != 1 || event.MouseButtonEvent.State != sdl.RELEASED {
		return nil
	}

	return &Command{
		TypeOf: CanvasClick,
		Target: sdl.Point{
			X: event.MouseButtonEvent.X,
			Y: event.MouseButtonEvent.Y,
		},
		TargetId: c.Id,
		Layer:    c.layer,
	}
}

func (c *Canvas) Update(cmd *Command) {

	switch cmd.TypeOf {
	case Noop:
	case CanvasClick:
		if cmd.TargetId != c.Id {
			return
		}

		c.curves[c.currentCurve].Update(cmd)
	case ControlPointPress, ControlPointRelease, MousePosition:
		if cmd.Target.InRect(c.dst) {
			c.curves[c.currentCurve].Update(cmd)
		}
	}
}

func (c *Canvas) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.FillRect(c.dst)

	for _, curve := range c.curves {
		curve.Render(renderer)
	}
}

func (c *Canvas) Rect() *sdl.Rect {
	return c.dst
}

func (c *Canvas) Register(colM map[int]Ui) {
	colM[c.Id] = c

	for _, cur := range c.curves {
		cur.Register(colM)
	}
}

func (c *Canvas) Layer() int {
	return c.layer
}
