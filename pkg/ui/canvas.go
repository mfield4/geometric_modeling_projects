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
		initCurveId: NewCasteljauBezierCurve(initCurveId, 2, true),
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
	if rect == nil {
		return
	}

	//renderer.SetDrawColor(255, 255, 255, 255)
	//renderer.DrawRect(rect)
	//
	//for _, pt := range c.Intersections() {
	//    gfx.CircleColor(renderer, pt.X, pt.Y, 5, sdl.Color{
	//        R: 0,
	//        G: 255,
	//        B: 255,
	//        A: 255,
	//    })
	//}

}

func (c *Canvas) Split(t float64) {
	curCurve := c.curves[c.currentCurve]

	l, r := curCurve.splitCurve(t, true)

	length := len(l.ctlPoints)

	l.ctlPoints[length-1].X -= 10
	l.ctlPoints[length-1].Y -= 10
	r.ctlPoints[0].X += 10
	r.ctlPoints[0].Y += 10

	l.Draw()
	r.Draw()

	c.curves[c.currentCurve] = l
	c.curves[r.Id] = r
}

func generateCombinations(input []*CasteljauBezierCurve, length int) <-chan []*CasteljauBezierCurve {
	c := make(chan []*CasteljauBezierCurve)

	go func(c chan []*CasteljauBezierCurve) {
		defer close(c)

		AddCombo(c, nil, input, length)
	}(c)

	return c
}

func AddCombo(c chan []*CasteljauBezierCurve, input, original []*CasteljauBezierCurve, length int) {
	if length <= 0 {
		return
	}

	var newCombo []*CasteljauBezierCurve
	for _, ch := range original {
		newCombo = append(input, ch)
		c <- newCombo
		AddCombo(c, newCombo, original, length-1)
	}
}

func curveIntersections(curve1 *CasteljauBezierCurve, curve2 *CasteljauBezierCurve) []sdl.Point {
	/*
	    c1l -> c2l
	    c1l -> c2r
	    c1r -> c2r
	    c1r -> c2r

	   Check If curves intersect.
	   If no, return nil.
	   If yes
	        Split curves. Find halves that intersect.
	        If Threshold reached, return midpoint of intersecting curves.

	        return curveIntersection(l,r)
	*/
	if !curve1.Rect().HasIntersection(curve2.Rect()) {
		return nil
	}

	c1l, c1r := curve1.splitCurve(0.5, false)
	c2l, c2r := curve2.splitCurve(0.5, false)

	var points []sdl.Point
	var curves []*CasteljauBezierCurve
	localCombs := generateCombinations(append(curves, c1l, c1r, c2l, c2r), 2)

	for cmb := range localCombs {
		if len(cmb) != 2 {
			continue
		}
		if cmb[0].Rect().HasIntersection(cmb[1].Rect()) {
			points = append(points, curveIntersections(cmb[0], cmb[1])...)
		}
	}

	return points
}

func (c *Canvas) Intersections() []sdl.Point {
	var intersections []sdl.Point
	var curves []*CasteljauBezierCurve
	if len(c.curves) < 2 {
		return nil
	}

	for _, cur := range c.curves {
		curves = append(curves, cur)
	}
	combinations := generateCombinations(curves, 2)

	for curveComb := range combinations {
		if len(curveComb) != 2 {
			continue
		}

		intersections = append(intersections, curveIntersections(curveComb[0], curveComb[1])...)
	}

	return intersections
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
