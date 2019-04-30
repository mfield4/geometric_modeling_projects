package canvas

import (
	"github.com/mfield4/178_projects/pkg/curves"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

// Basically same as main canvas. However, it instead raises and draws a raised curve.
type RaisedCanvas struct {
	mc             Canvas
	x0, y0         int32
	xScale, yScale float32
	dst            *sdl.Rect
}

func NewRaisedCanvas(canvas Canvas) *RaisedCanvas {
	return &RaisedCanvas{
		mc:     canvas,
		x0:     canvas.Rect().X,
		y0:     canvas.Rect().H,
		xScale: 0.5,
		yScale: 1,
		dst: &sdl.Rect{
			X: canvas.Rect().X,
			Y: canvas.Rect().H,
			W: canvas.Rect().W / 2,
			H: canvas.Rect().H,
		},
	}
}

func (rc *RaisedCanvas) Render(renderer *sdl.Renderer) {
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

	err = renderer.SetDrawColor(66, 98, 244, 255)
	err = renderer.DrawLines(rc.Scale(points...))
	err = renderer.SetDrawColor(0, 0, 0, 255)
	err = renderer.DrawLines(rc.Scale(ctl...))

	if err != nil {
		log.Println(err)
	}
}

func (rc *RaisedCanvas) Scale(pts ...sdl.Point) []sdl.Point {
	scaled := make([]sdl.Point, len(pts))
	for i, pt := range pts {
		scaled[i].X = int32(float32(pt.X)*rc.xScale) + rc.x0
		scaled[i].Y = int32(float32(pt.Y)*rc.yScale) + rc.y0
	}

	return scaled
}

func (*RaisedCanvas) Update() {
	panic("implement me")
}

func (rc *RaisedCanvas) Curve() curves.Curve {
	origC := rc.mc.Curve()
	n := int32(len(origC.Poly()))
	if n < 3 {
		return origC
	}

	orig := func(i int32) sdl.Point {
		if i == -1 || i >= n {
			return sdl.Point{}
		}
		return origC.Poly()[i]
	}

	newPts := make([]sdl.Point, n+1)

	var i int32
	for i = 0; i < n+1; i++ {
		coef1 := float32(i) / float32(n)
		coef2 := float32(n-i) / float32(n)
		newPts[i].X = int32(coef1*float32(orig(i-1).X) + coef2*float32(orig(i).X))
		newPts[i].Y = int32(coef1*float32(orig(i-1).Y) + coef2*float32(orig(i).Y))
	}

	newCurve := curves.NewBezierCurve(origC.Id(), 1, false)
	newCurve.Add(newPts...)

	return newCurve
}

func (rc *RaisedCanvas) Rect() sdl.Rect {
	return *rc.dst
}
