package canvas

import (
	"github.com/mfield4/178_projects/pkg/curves"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

// Basically same as main canvas. However, it instead raises and draws a raised curve.
type ReducedCanvas struct {
	mc             Canvas
	x0, y0         int32
	xScale, yScale float32
	dst            *sdl.Rect
}

func NewReducedCanvas(canvas Canvas) *ReducedCanvas {
	return &ReducedCanvas{
		mc:     canvas,
		x0:     canvas.Rect().W / 2,
		y0:     canvas.Rect().H,
		xScale: 0.5,
		yScale: 1,
		dst: &sdl.Rect{
			X: canvas.Rect().W / 2,
			Y: canvas.Rect().H,
			W: canvas.Rect().W / 2,
			H: canvas.Rect().H,
		},
	}
}

func (rc *ReducedCanvas) Render(renderer *sdl.Renderer) {
	curve := rc.Curve()

	err := renderer.SetDrawColor(255, 255, 255, 255)
	err = renderer.FillRect(rc.dst)
	if err != nil {
		log.Println(err)
		return
	}

	drawPoints := func(ctl []sdl.Point) {
		for _, pt := range rc.Scale(ctl...) {
			gfx.CircleColor(renderer, pt.X, pt.Y, 5, sdl.Color{
				R: 255,
				G: 0,
				B: 0,
				A: 255,
			})
		}
	}

	ctl := curve.Poly()
	drawPoints(ctl)

	if len(ctl) < 2 {
		return
	}

	points := curve.Curve()

	err = renderer.SetDrawColor(244, 78, 66, 255)
	err = renderer.DrawLines(rc.Scale(points...))
	err = renderer.SetDrawColor(0, 0, 0, 255)
	err = renderer.DrawLines(rc.Scale(ctl...))

	if err != nil {
		log.Println(err)
	}
}

func (rc *ReducedCanvas) Scale(pts ...sdl.Point) []sdl.Point {
	scaled := make([]sdl.Point, len(pts))
	for i, pt := range pts {
		scaled[i].X = int32(float32(pt.X)*rc.xScale) + rc.x0
		scaled[i].Y = int32(float32(pt.Y)*rc.yScale) + rc.y0
	}

	return scaled
}

func (*ReducedCanvas) Update() {
	panic("implement me")
}

func (rc *ReducedCanvas) Curve() curves.Curve {
	origC := rc.mc.Curve()
	n := int32(len(origC.Poly()))
	if n < 3 {
		return origC
	}

	index := func(list []sdl.Point, i int32) sdl.Point {
		if i == -1 || i >= int32(len(list)) {
			return sdl.Point{}
		}
		return list[i]
	}

	left := func(pts []sdl.Point) []sdl.Point {
		newPts := make([]sdl.Point, n-1)

		var i int32
		for ; i < n-1; i++ {
			coef1 := float64(n) / float64(n-i)
			coef2 := float64(i) / float64(n-i)
			newPts[i].X = int32(coef1*float64(index(pts, i).X) - coef2*float64(index(newPts, i-1).X))
			newPts[i].Y = int32(coef1*float64(index(pts, i).Y) - coef2*float64(index(newPts, i-1).Y))
		}

		return newPts
	}

	right := func(pts []sdl.Point) []sdl.Point {
		newPts := make([]sdl.Point, n-1)

		var i int32
		for i = n - 1; i > 0; i-- {
			coef1 := float64(n-1) / float64(i)
			coef2 := float64(n-1-i) / float64(i)
			newPts[i-1].X = int32(coef1*float64(index(pts, i).X) - coef2*float64(index(newPts, i).X))
			newPts[i-1].Y = int32(coef1*float64(index(pts, i).Y) - coef2*float64(index(newPts, i).Y))
		}

		return newPts
	}

	newCurve := curves.NewBezierCurve(origC.Id(), 1, false)
	l, r := left(origC.Poly()), right(origC.Poly())
	avg := MidpointAvg(l, r)
	newCurve.Add(avg...)

	return newCurve
}

func Midpoint(pt1, pt2 sdl.Point) sdl.Point {
	return sdl.Point{
		X: (pt2.X + pt1.X) / 2,
		Y: (pt2.Y + pt1.Y) / 2,
	}
}

func MidpointAvg(left, right []sdl.Point) (avg []sdl.Point) {
	n := len(left)

	avg = make([]sdl.Point, n)

	for i := 0; i < n; i++ {
		avg[i] = Midpoint(left[i], right[i])
	}
	return
}

func (rc *ReducedCanvas) Rect() sdl.Rect {
	return *rc.dst
}
