package curves

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

var Steps = 1000
var Bern = true

type CasteljauBezierCurve struct {
	id    int
	layer int

	focused bool
	index   int

	ctlPoints   []sdl.Point
	curvePoints []sdl.Point
	bernstein   [][]func(t float64) float64
}

func NewBezierCurve(ID, layer int, register bool) *CasteljauBezierCurve {
	cbc := &CasteljauBezierCurve{
		id:          ID,
		layer:       layer,
		focused:     false,
		index:       0,
		ctlPoints:   make([]sdl.Point, 0),
		curvePoints: make([]sdl.Point, Steps),
		bernstein:   make([][]func(float64) float64, 0),
	}

	return cbc
}

func (cbc *CasteljauBezierCurve) Id() int {
	return cbc.id
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

	cbc.Draw()
}

func (cbc *CasteljauBezierCurve) Curve() []sdl.Point {
	return cbc.curvePoints
}

func (cbc *CasteljauBezierCurve) Poly() []sdl.Point {
	return cbc.ctlPoints
}

func (cbc *CasteljauBezierCurve) Press(state, x, y int32) {
	if state == 0 {
		cbc.focused = false
		cbc.Draw()
		return
	}

	if cbc.focused {
		return
	}

	cbc.focused = true

	mousePt := sdl.Point{
		X: x,
		Y: y,
	}

	low := math.MaxFloat64
	for i, pt := range cbc.ctlPoints {
		cur := Dist(pt, mousePt)
		if cur < low {
			cbc.index = i
			low = cur
		}
	}

	cbc.ctlPoints[cbc.index].X, cbc.ctlPoints[cbc.index].Y = x, y
	fmt.Printf("PRESSED x: %d, y %d\n", cbc.ctlPoints[cbc.index].X, cbc.ctlPoints[cbc.index].Y)
	cbc.Draw()
}

func (cbc *CasteljauBezierCurve) Drag(x, y int32) {
	if !cbc.focused {
		return
	}

	cbc.ctlPoints[cbc.index].X = x
	cbc.ctlPoints[cbc.index].Y = y
	cbc.Draw()
}

var THRESHOLD = 10.0

func (cbc *CasteljauBezierCurve) Layer(state, x, y int32) int {
	if state != 1 {
		if cbc.focused {
			return cbc.layer
		}

		return 0
	}

	mousePt := sdl.Point{
		X: x,
		Y: y,
	}

	low := math.MaxFloat64
	for _, pt := range cbc.ctlPoints {
		cur := Dist(pt, mousePt)
		if cur < low {
			low = cur
		}
	}

	if low > THRESHOLD {
		return 0
	}

	return cbc.layer
}

func Dist(pt1, pt2 sdl.Point) float64 {
	return math.Sqrt(math.Pow(float64(pt2.X-pt1.X), 2) - math.Pow(float64(pt2.Y-pt2.Y), 2))
}

func round(val float64) int32 {
	if val < 0 {
		return int32(val - 0.5)
	}
	return int32(val + 0.5)
}
