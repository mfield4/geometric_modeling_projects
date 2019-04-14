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
