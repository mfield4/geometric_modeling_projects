package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Canvas struct {
	Id           int
	layer        int
	currentCurve int
	curves       map[int]*CasteljauBezierCurve

	dst *sdl.Rect
	pc  *PreviewCanvas
}

func NewCanvas(width int32, height int32) *Canvas {
	initCurveId := GUID()
	castel := map[int]*CasteljauBezierCurve{
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

func (c *Canvas) PressActive(x, y int32) bool {
	pt := sdl.Point{X: x, Y: y}

	if pt.InRect(c.dst) {
		return true
	}

	return false
}

func (c *Canvas) Mouse1Down(x, y int32) {
	c.curves[c.currentCurve].Add(sdl.Point{
		X: x,
		Y: y,
	})
	c.curves[c.currentCurve].Draw()
	c.pc.Draw()
}

func (c *Canvas) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.FillRect(c.dst)

	for _, curve := range c.curves {
		curve.Render(renderer)
	}

	c.pc.Render(renderer)
}

func (c *Canvas) RegisterCol(colM map[int]Ui) {
	colM[c.Id] = c

	for _, cur := range c.curves {
		cur.RegisterCol(colM)
	}
}

func (c *Canvas) RegisterM1d(downs map[int]Mouse1Down) {
	downs[c.Id] = c

	for _, cur := range c.curves {
		cur.RegisterM1d(downs)
	}
}

func (c *Canvas) RegisterM1u(up map[int]Mouse1Up) {
	for _, cur := range c.curves {
		cur.RegisterM1u(up)
	}
}

func (c *Canvas) RegisterMM(mm map[int]MouseMotion) {
	for _, cur := range c.curves {
		cur.RegisterMM(mm)
	}
}

func (c *Canvas) Layer() int {
	return c.layer
}
