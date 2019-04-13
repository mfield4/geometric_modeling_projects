package ui

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type BezierCurve interface {
	Ui
	Add(...sdl.Point)
	Draw()
}

var Steps = 1000

type CasteljauBezierCurve struct {
	Id          int
	ctlPoints   []sdl.Point
	curvePoints []sdl.Point
}

func NewCasteljauBezierCurve(ID int) *CasteljauBezierCurve {
	return &CasteljauBezierCurve{
		Id:          ID,
		ctlPoints:   make([]sdl.Point, 0),
		curvePoints: make([]sdl.Point, Steps),
	}
}

func (cbc *CasteljauBezierCurve) Add(points ...sdl.Point) {
	cbc.ctlPoints = append(cbc.ctlPoints, points...)
}

func (cbc *CasteljauBezierCurve) Draw() {
	// Do the algorithm to get the points from the control points
}

func (cbc *CasteljauBezierCurve) Input(*sdl.MouseButtonEvent) *Command {
	panic("implement me")
}

func (cbc *CasteljauBezierCurve) Update(*Command) {
	panic("implement me")
}

func (cbc *CasteljauBezierCurve) Render(renderer *sdl.Renderer) {
	if len(cbc.ctlPoints) == 0 {
		return
	}

	for _, point := range cbc.ctlPoints {
		gfx.CircleColor(renderer, point.X, point.Y, 2, sdl.Color{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
	}

	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawLines(cbc.ctlPoints)

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLines(cbc.curvePoints)
}

func (cbc *CasteljauBezierCurve) Rect() *sdl.Rect {
	panic("implement me")
}

func (cbc *CasteljauBezierCurve) Register(map[int]Ui) {
	panic("implement me")
}

func (cbc *CasteljauBezierCurve) Layer() int {
	panic("implement me")
}
