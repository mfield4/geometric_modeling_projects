package canvas

import (
	"github.com/mfield4/178_projects/pkg/curves"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

//
//import (
//	"github.com/veandco/go-sdl2/gfx"
//	"github.com/veandco/go-sdl2/sdl"
//)
//
//var THRESHOLD int32 = 50
//var SplitVal = 0.5
////

type MainCanvas struct {
	Id           int
	layer        int
	currentCurve int
	curves       map[int]curves.Curve

	dst *sdl.Rect
}

func NewMainCanvas(x, y, width, height int32, initCurve curves.Curve) *MainCanvas {
	mainCanvas := &MainCanvas{
		Id:           0,
		layer:        1,
		currentCurve: initCurve.Id(),
		curves: map[int]curves.Curve{
			initCurve.Id(): initCurve,
		},
		dst: &sdl.Rect{
			X: x,
			Y: y,
			W: width,
			H: height,
		},
	}

	return mainCanvas
}

func (mc *MainCanvas) Press(button, x, y int32) {
	mc.curves[mc.currentCurve].Add(sdl.Point{X: x, Y: y})
}

func (mc *MainCanvas) Layer(button, x, y int32) int {
	pt := sdl.Point{
		X: x,
		Y: y,
	}

	if button != 1 || !pt.InRect(mc.dst) {
		return 0
	}

	return mc.layer
}

func (mc *MainCanvas) Render(renderer *sdl.Renderer) {
	err := renderer.SetDrawColor(255, 255, 255, 255)
	err = renderer.FillRect(mc.dst)
	if err != nil {
		log.Println(err)
		return
	}

	drawPoints := func(ctl []sdl.Point) {
		for _, pt := range ctl {
			gfx.CircleColor(renderer, pt.X, pt.Y, 5, sdl.Color{
				R: 255,
				G: 0,
				B: 0,
				A: 255,
			})
		}

	}

	for _, curve := range mc.curves {
		ctl := curve.Poly()
		drawPoints(ctl)

		if len(ctl) < 2 {
			continue
		}

		points := curve.Curve()

		err := renderer.SetDrawColor(0, 255, 255, 255)
		err = renderer.DrawLines(points)
		err = renderer.SetDrawColor(0, 0, 0, 255)
		err = renderer.DrawLines(ctl)

		if err != nil {
			log.Println(err)
		}
	}
}

func (*MainCanvas) Update() {
	panic("implement me")
}

func (mc *MainCanvas) Ref() Canvas {
	return mc
}

//
//func NewCanvas(width int32, height int32) *Canvas {
//	initCurveId := GUID()
//	castel := map[int]*CasteljauBezierCurve{
//		initCurveId: NewCasteljauBezierCurve(initCurveId, 2, true),
//	}
//
//	canvas := &Canvas{
//		Id:           GUID(),
//		layer:        1,
//		currentCurve: initCurveId,
//		curves:       castel,
//		dst: &sdl.Rect{
//			X: 0,
//			Y: 0,
//			W: width,
//			H: height,
//		},
//	}
//	canvas.RegisterCol()
//	canvas.RegisterM1d()
//	canvas.RegisterM2d()
//	return canvas
//}
//
//func (c *Canvas) Render(renderer *sdl.Renderer) {

//
//}
//
//func (c *Canvas) Split(t float64) {
//	curCurve := c.curves[c.currentCurve]
//
//	l, r := curCurve.splitCurve(t, true)
//
//	length := len(l.ctlPoints)
//
//	l.ctlPoints[length-1].X -= 10
//	l.ctlPoints[length-1].Y -= 10
//	r.ctlPoints[0].X += 10
//	r.ctlPoints[0].Y += 10
//
//	l.Draw()
//	r.Draw()
//
//	c.curves[c.currentCurve] = l
//	c.curves[r.Id] = r
//}
//
//func generateCombinations(input []*CasteljauBezierCurve, length int) (combos [][]*CasteljauBezierCurve) {
//	n := len(input)
//	for num := 0; num < (1 << uint(n)); num++ {
//		combination := []*CasteljauBezierCurve{}
//		for ndx := 0; ndx < n; ndx++ {
//			// (is the bit "on" in this number?)
//			if num&(1<<uint(ndx)) != 0 && input[ndx] != nil {
//				// (then add it to the combination)
//				combination = append(combination, input[ndx])
//			}
//		}
//		//fmt.Println(combination)
//		if len(combination) == 2 {
//			combos = append(combos, combination)
//		}
//	}
//
//	return combos
//}
//
//func AddCombo(c chan []*CasteljauBezierCurve, input, original []*CasteljauBezierCurve, length int) {
//	if length <= 0 {
//		return
//	}
//
//	var newCombo []*CasteljauBezierCurve
//	for _, ch := range original {
//		newCombo = append(input, ch)
//		c <- newCombo
//		AddCombo(c, newCombo, original, length-1)
//	}
//}
//
//func curveIntersections(curve1 *CasteljauBezierCurve, curve2 *CasteljauBezierCurve) (points []sdl.Point) {
//	/*
//	    c1l -> c2l
//	    c1l -> c2r
//	    c1r -> c2r
//	    c1r -> c2r
//
//	   Check If curves intersect.
//	   If no, return nil.
//	   If yes
//	        Split curves. Find halves that intersect.
//	        If Threshold reached, return midpoint of intersecting curves.
//
//	        return curveIntersection(l,r)
//	*/
//	if !curve1.Rect().HasIntersection(curve2.Rect()) {
//		return nil
//	}
//
//	rect1 := curve1.Rect()
//	rect2 := curve2.Rect()
//	if rectArea(rect1) < THRESHOLD || rectArea(rect2) < THRESHOLD {
//		points = append(points, midpoint(rectMidpoint(rect1), rectMidpoint(rect2)))
//		return
//	}
//
//	c1l, c1r := curve1.splitCurve(0.5, false)
//	c2l, c2r := curve2.splitCurve(0.5, false)
//
//	//var curves []*CasteljauBezierCurve
//	//localCombs := generateCombinations(append(curves, c1l, c1r, c2l, c2r), 2)
//	localCombs := [][]*CasteljauBezierCurve{
//		{c1l, c2l},
//		{c1l, c2r},
//		{c1r, c2l},
//		{c1r, c2r},
//	}
//
//	for _, cmb := range localCombs {
//		if len(cmb) != 2 {
//			continue
//		}
//		r1 := cmb[0].Rect()
//		r2 := cmb[1].Rect()
//
//		_, ok := r1.Intersect(r2)
//		if !ok {
//			continue
//		}
//		//if ri.W*ri.H < THRESHOLD/5 {
//		//	continue
//		//}
//
//		r1Size := r1.W * r1.W
//		r2Size := r2.W * r2.H
//		if r1Size < THRESHOLD || r2Size < THRESHOLD {
//			points = append(points, midpoint(rectMidpoint(r1), rectMidpoint(r2)))
//			break
//		}
//		newPoints := curveIntersections(cmb[0], cmb[1])
//		if newPoints != nil {
//			points = append(points, newPoints[0])
//		}
//	}
//
//	return points
//}
//
//func rectArea(rect *sdl.Rect) int32 {
//	return rect.W * rect.H
//}
//
//func rectMidpoint(r *sdl.Rect) sdl.Point {
//	x := r.X + r.W/2
//	y := r.Y + r.H/2
//	return sdl.Point{X: x, Y: y}
//}
//
//func midpoint(a, b sdl.Point) sdl.Point {
//	x, y := (a.X+b.X)/2, (a.Y+b.Y)/2
//	return sdl.Point{X: x, Y: y}
//}
//
//func (c *Canvas) Intersections() []sdl.Point {
//	var intersections []sdl.Point
//	var curves []*CasteljauBezierCurve
//	if len(c.curves) < 2 {
//		return nil
//	}
//
//	for _, cur := range c.curves {
//		curves = append(curves, cur)
//	}
//	combinations := generateCombinations(curves, 2)
//
//	for _, curveComb := range combinations {
//		if len(curveComb) != 2 {
//			continue
//		}
//
//		intersections = append(intersections, curveIntersections(curveComb[0], curveComb[1])...)
//	}
//
//	return intersections
//}
//
//func (c *Canvas) PressActive(x, y int32) bool {
//	pt := sdl.Point{X: x, Y: y}
//
//	if pt.InRect(c.dst) {
//		return true
//	}
//
//	return false
//}
//
//func (c *Canvas) Mouse1Down(x, y int32) {
//	c.curves[c.currentCurve].Add(sdl.Point{
//		X: x,
//		Y: y,
//	})
//	c.curves[c.currentCurve].Draw()
//}
//
//func (c *Canvas) Mouse2Down(x, y int32) {
//	mousePt := sdl.Point{
//		X: x,
//		Y: y,
//	}
//	if mousePt.InRect(c.dst) {
//		c.Split(SplitVal)
//	}
//}
//
//func (c *Canvas) RightActive() bool {
//	return true
//}
//
//func (c *Canvas) RegisterCol() {
//	GetCollisionMap()[c.Id] = c
//}
//
//func (c *Canvas) RegisterM1d() {
//	GetMouse1dCommands()[c.Id] = c
//}
//
//func (c *Canvas) RegisterM2d() {
//	GetMouse2dCommands()[c.Id] = c
//}
//
//func (c *Canvas) Layer() int {
//	return c.layer
//}
