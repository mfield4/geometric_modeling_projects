package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/big"
)

func (cbc *CasteljauBezierCurve) Draw() {
	// Do the algorithm to get the points from the control points
	if cbc.bern {
		cbc.drawBern()
		return
	}

	cbc.drawCastel()
}

func (cbc *CasteljauBezierCurve) drawCastel() {
	if len(cbc.ctlPoints) < 3 {
		return
	}
	for i := 0; i < Steps; i++ {
		p := cbc.casteljauCurvePoint(float64(i) / float64(Steps))
		cbc.curvePoints[i] = p
	}
}

func (cbc *CasteljauBezierCurve) casteljauCurvePoint(t float64) sdl.Point {
	lPoints, _ := cbc.splitCurve(t)

	return lPoints[len(cbc.ctlPoints)-1]
}

func (cbc *CasteljauBezierCurve) splitCurve(t float64) (l []sdl.Point, r []sdl.Point) {
	length := len(cbc.ctlPoints)
	ctlPoints := make([][]sdl.Point, length+1)
	l = make([]sdl.Point, length)
	r = make([]sdl.Point, length)

	for i := 0; i < length+1; i++ {
		ctlPoints[i] = make([]sdl.Point, length+1)
	}
	for i := range cbc.ctlPoints {
		ctlPoints[0][i] = cbc.ctlPoints[i]
	}
	for j := 1; j < length; j++ {
		for i := 0; i < (length - j); i++ {
			pointFunc := func(pt1, pt2 sdl.Point) sdl.Point {
				pt := sdl.Point{
					X: round((1-t)*float64(pt1.X) + t*float64(pt2.X)),
					Y: round((1-t)*float64(pt1.Y) + t*float64(pt2.Y)),
				}
				return pt
			}
			ctlPoints[j][i] = pointFunc(ctlPoints[j-1][i], ctlPoints[j-1][i+1])
		}
	}

	for k := 0; k < length; k++ {
		l[k].X, l[k].Y = ctlPoints[k][0].X, ctlPoints[k][0].Y
		r[k].X, r[k].Y = ctlPoints[length-k-1][k].X, ctlPoints[length-k-1][k].Y
	}

	return
}

func (cbc *CasteljauBezierCurve) drawBern() {
	if len(cbc.ctlPoints) < 3 {
		return
	}

	for t := 0; t < Steps; t++ {
		toAdd := sdl.Point{}
		var x, y float64
		for i, pt := range cbc.ctlPoints {
			var nt float64 = float64(t) / float64(Steps)
			if math.Abs(nt-1) < 0.01 {
				fmt.Printf("%+v\n", nt)
			}
			nt = cbc.bernstein[len(cbc.ctlPoints)-1][i](nt)
			x += float64(pt.X) * nt
			y += float64(pt.Y) * nt
		}
		toAdd.X, toAdd.Y = int32(x), int32(y)
		cbc.curvePoints[t] = toAdd
	}
}

func bernsteinPoly(k, n int64) func(t float64) float64 {
	binom := big.NewInt(0)
	binom = binom.Binomial(n, k)
	return func(t float64) float64 {
		var nt float64 = math.Pow(float64(t), float64(k)) * math.Pow(float64(1-t), float64(n-k))
		return float64(binom.Int64()) * nt
	}
}
