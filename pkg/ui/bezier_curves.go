package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

var Steps = 1000
var Bern = true

type BezierCurve interface {
	Ui
	Add(...sdl.Point)
	Draw()
	RegisterM1d(downs map[int]Mouse1Down)
	RegisterM1u(ups map[int]Mouse1Up)
	RegisterMM(mouseMotions map[int]MouseMotion)
}

type CasteljauBezierCurve struct {
	Id    int
	layer int

	focused bool
	index   int

	ctlPoints   []sdl.Point
	curvePoints []sdl.Point
	bernstein   [][]func(t float64) float64
}

func NewCasteljauBezierCurve(ID, layer int, register bool) *CasteljauBezierCurve {
	cbc := &CasteljauBezierCurve{
		Id:          ID,
		layer:       layer,
		focused:     false,
		index:       0,
		ctlPoints:   make([]sdl.Point, 0),
		curvePoints: make([]sdl.Point, Steps),
		bernstein:   make([][]func(float64) float64, 0),
	}

	if register {
		cbc.RegisterCol()
		cbc.RegisterM1d()
		cbc.RegisterM1u()
		cbc.RegisterMM()
	}
	return cbc
}

func (cbc *CasteljauBezierCurve) Add(points ...sdl.Point) {
	for _, pt := range points {
		cbc.ctlPoints = append(cbc.ctlPoints, pt)
	}

	numPts := len(cbc.ctlPoints)
	cbc.bernstein = make([][]func(float64) float64, numPts)
	for i := 0; i < numPts; i++ {
		cbc.bernstein[i] = make([]func(float64) float64, numPts)
		for j := 0; j < i+1; j++ {
			cbc.bernstein[i][j] = bernsteinPoly(int64(j), int64(i))
		}
	}
}

func (cbc *CasteljauBezierCurve) PressActive(x, y int32) bool {
	cmp := func(p1, p2 sdl.Point) float64 {
		first := math.Pow(float64(p2.X-p1.X), 2)
		second := math.Pow(float64(p2.Y-p1.Y), 2)
		return math.Sqrt(first + second)
	}

	mousePt := sdl.Point{
		X: x,
		Y: y,
	}

	for i, pt := range cbc.ctlPoints {
		if cmp(mousePt, pt) < 20 {
			cbc.index = i
			return true
		}
	}

	return false
}

func (cbc *CasteljauBezierCurve) Mouse1Down(x, y int32) {
	if cbc.focused {
		return
	}
	cbc.focused = true

	cbc.ctlPoints[cbc.index].X, cbc.ctlPoints[cbc.index].Y = x, y
	fmt.Printf("PRESSED x: %d, y %d\n", cbc.ctlPoints[cbc.index].X, cbc.ctlPoints[cbc.index].Y)
	cbc.Draw()
}

func (cbc *CasteljauBezierCurve) ReleaseActive(x, y int32) bool {
	return cbc.focused
}

func (cbc *CasteljauBezierCurve) Mouse1Up(x, y int32) {
	cbc.focused = false
	cbc.Draw()

}

func (cbc *CasteljauBezierCurve) MotionActive() bool {
	return cbc.focused
}

func (cbc *CasteljauBezierCurve) MouseMotion(x, y int32) {
	cbc.ctlPoints[cbc.index].X = x
	cbc.ctlPoints[cbc.index].Y = y
	cbc.Draw()
}

func (cbc *CasteljauBezierCurve) RegisterCol() {
	GetCollisionMap()[cbc.Id] = cbc
}

func (cbc *CasteljauBezierCurve) RegisterM1d() {
	GetMouse1dCommands()[cbc.Id] = cbc
}

func (cbc *CasteljauBezierCurve) RegisterM1u() {
	GetMouse1uCommands()[cbc.Id] = cbc
}

func (cbc *CasteljauBezierCurve) RegisterMM() {
	GetMouseMMCommands()[cbc.Id] = cbc
}

func (cbc *CasteljauBezierCurve) Layer() int {
	return cbc.layer
}

func (cbc *CasteljauBezierCurve) Render(renderer *sdl.Renderer) {
	if len(cbc.ctlPoints) == 0 {
		return
	}

	for _, point := range cbc.ctlPoints {
		gfx.CircleColor(renderer, point.X, point.Y, 5, sdl.Color{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
	}

	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawLines(cbc.ctlPoints)

	if len(cbc.curvePoints) == 0 {
		return
	}

	//if cbc.focused {
	//	renderer.SetDrawColor(255, 255, 255, 255/2)
	//	renderer.DrawRect(cbc.Rect())
	//}
	if Bern {
		renderer.SetDrawColor(0, 255, 255, 255)
	} else {
		renderer.SetDrawColor(255, 0, 255, 255)
	}
	renderer.DrawLines(cbc.curvePoints)
}

func (cbc *CasteljauBezierCurve) Rect() *sdl.Rect {
	if len(cbc.ctlPoints) == 0 {
		return nil
	}

	canvasW := (WindowWidth * 2) / 3

	dst := &sdl.Rect{
		X: 0,
		Y: 0,
		W: canvasW,
		H: WindowHeight,
	}

	minBoundedRect, ok := sdl.EnclosePoints(cbc.ctlPoints, dst)
	if !ok {
		return nil
	}

	return &minBoundedRect
}

func (cbc *CasteljauBezierCurve) current() []sdl.Point {
	return cbc.curvePoints
}

func round(val float64) int32 {
	if val < 0 {
		return int32(val - 0.5)
	}
	return int32(val + 0.5)
}
