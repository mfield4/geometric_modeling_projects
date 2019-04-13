package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Canvas struct {
	ID           int
	layer        int
	currentCurve int
	curves       map[int]BezierCurve

	dst *sdl.Rect
}

func NewCanvas(width int32, height int32) *Canvas {
	initCurveId := GUID()
	curves := map[int]BezierCurve{
		initCurveId: NewCasteljauBezierCurve(initCurveId),
	}

	return &Canvas{
		ID:           GUID(),
		layer:        1,
		currentCurve: initCurveId,
		curves:       curves,
		dst: &sdl.Rect{
			X: 0,
			Y: 0,
			W: width,
			H: height,
		},
	}
}

func (c *Canvas) Input(event *sdl.MouseButtonEvent) *Command {
	// Exit conditions. Only want a left click release
	if event.Button != 1 || event.State != sdl.RELEASED {
		return nil
	}

	return &Command{
		typeOf: CanvasClick,
		target: sdl.Point{
			X: event.X,
			Y: event.Y,
		},
		targetId: c.ID,
	}
}

func (c *Canvas) Update(cmd *Command) {
	if cmd.targetId != c.ID {
		return
	}

	switch cmd.typeOf {
	case Noop:
	case CanvasClick:
		c.addToCurves(cmd.target)
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
	colM[c.ID] = c
}

func (c *Canvas) Layer() int {
	return c.layer
}

func (c *Canvas) addToCurves(point sdl.Point) {
	c.curves[c.currentCurve].Add(point)
}
