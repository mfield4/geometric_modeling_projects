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
}

func NewCanvas(width int32, height int32) *Canvas {
	initCurveId := GUID()
	castel := map[int]*CasteljauBezierCurve{
		initCurveId: NewCasteljauBezierCurve(initCurveId, 2),
	}

	canvas := &Canvas{
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
	canvas.RegisterCol()
	canvas.RegisterM1d()
	canvas.RegisterM2d()
	return canvas
}

func (c *Canvas) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.FillRect(c.dst)

	for _, curve := range c.curves {
		curve.Render(renderer)
		if curve.focused {
			c.currentCurve = curve.Id
		}
	}

	rect := c.curves[c.currentCurve].Rect()
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawRect(rect)
}

func (c *Canvas) Split(t float64) {
	curCurve := c.curves[c.currentCurve]

	newCurveL := NewCasteljauBezierCurve(c.currentCurve, 2)
	newCurveR := NewCasteljauBezierCurve(GUID(), 2)

	l, r := curCurve.splitCurve(t)

	length := len(l)

	l[length-1].X -= 10
	l[length-1].Y -= 10
	r[0].X += 10
	r[0].Y += 10

	newCurveL.Add(l...)
	newCurveR.Add(r...)

	newCurveL.Draw()
	newCurveR.Draw()

	c.curves[c.currentCurve] = newCurveL
	c.curves[newCurveR.Id] = newCurveR
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
}

func (c *Canvas) Mouse2Down(x, y int32) {
	mousePt := sdl.Point{
		X: x,
		Y: y,
	}
	if mousePt.InRect(c.dst) {
		c.Split(0.5)
	}
}

func (c *Canvas) RightActive() bool {
	return true
}

func (c *Canvas) RegisterCol() {
	GetCollisionMap()[c.Id] = c
}

func (c *Canvas) RegisterM1d() {
	GetMouse1dCommands()[c.Id] = c
}

func (c *Canvas) RegisterM2d() {
	GetMouse2dCommands()[c.Id] = c
}

func (c *Canvas) Layer() int {
	return c.layer
}
