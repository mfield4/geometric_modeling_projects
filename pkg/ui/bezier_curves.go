package ui

import (
	"fmt"
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
	Id    int
	layer int

	focused bool
	index   int

	ctlPoints   []sdl.Point
	curvePoints []sdl.Point
}

func (cbc *CasteljauBezierCurve) Selected(s bool) {
	cbc.focused = s
}

func (cbc *CasteljauBezierCurve) IsSelected() bool {
	return cbc.focused
}

func NewCasteljauBezierCurve(ID, layer int) *CasteljauBezierCurve {
	return &CasteljauBezierCurve{
		Id:          ID,
		layer:       layer,
		focused:     false,
		index:       0,
		ctlPoints:   make([]sdl.Point, 0),
		curvePoints: make([]sdl.Point, Steps),
	}
}

func (cbc *CasteljauBezierCurve) Add(points ...sdl.Point) {
	cbc.ctlPoints = append(cbc.ctlPoints, points...)
}

func (cbc *CasteljauBezierCurve) Draw() {
	// Do the algorithm to get the points from the control points
	if len(cbc.ctlPoints) < 3 {
		return
	}
	for i := 0; i < Steps; i++ {
		p := cbc.casteljauCurvePoint(float64(i) / float64(Steps))
		cbc.curvePoints[i] = p
	}

}

func (cbc *CasteljauBezierCurve) casteljauCurvePoint(t float64) sdl.Point {
	length := len(cbc.ctlPoints)
	ctlPoints := make([]sdl.Point, length+1)

	for i := range cbc.ctlPoints {
		ctlPoints[i] = cbc.ctlPoints[i]
	}

	for k := 1; k < length; k++ {
		for i := 0; i < (length - k); i++ {
			pointFunc := func(pt1, pt2 sdl.Point) sdl.Point {
				return sdl.Point{
					X: round((1-t)*float64(pt1.X) + t*float64(pt2.X)),
					Y: round((1-t)*float64(pt1.Y) + t*float64(pt2.Y)),
				}
			}
			ctlPoints[i] = pointFunc(ctlPoints[i], ctlPoints[i+1])

		}
	}

	return ctlPoints[0]
}

func (cbc *CasteljauBezierCurve) Input(event *MouseEvent) *Command {
	if cbc.focused {
		return cbc.inputRelease(event)
	}
	return cbc.inputPress(event)

}

func (cbc *CasteljauBezierCurve) Update(cmd *Command) {
	switch cmd.TypeOf {
	case ControlPointPress:
		cbc.cntlPointPressUpdate(cmd)

	case ControlPointRelease:
		cbc.cntlPointRelease(cmd)

	case MousePosition:
		cbc.mousePositionCommand(cmd)

	case CanvasClick:
		cbc.canvasClickCommand(cmd)

	default:
		return
	}
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

	if cbc.focused {
		renderer.SetDrawColor(255, 255, 255, 255/2)
		renderer.DrawRect(cbc.Rect())
	}
	renderer.SetDrawColor(0, 0, 255, 255)
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

func (cbc *CasteljauBezierCurve) Register(colM map[int]Ui) {
	colM[cbc.Id] = cbc
}

func (cbc *CasteljauBezierCurve) Layer() int {
	return cbc.layer
}

func (cbc *CasteljauBezierCurve) inputPress(event *MouseEvent) *Command {
	if event.Button != 1 || event.MouseButtonEvent.State != sdl.PRESSED {
		return nil
	}

	mousePoint := sdl.Point{
		X: event.MouseButtonEvent.X,
		Y: event.MouseButtonEvent.Y,
	}

	checkPoint := func() (int, bool) {
		for i, pts := range cbc.ctlPoints {
			points := []sdl.Point{mousePoint, pts}
			rect, ok := sdl.EnclosePoints(points, nil)
			if !ok {
				continue
			}
			if rect.W > 15 || rect.H > 15 {
				continue
			}

			return i, true
		}
		return 0, false
	}

	id, ok := checkPoint()
	if !ok {
		return nil
	}

	return &Command{
		TypeOf:   ControlPointPress,
		Target:   mousePoint,
		TargetId: id,
		Layer:    cbc.layer,
	}
}

func (cbc *CasteljauBezierCurve) inputRelease(event *MouseEvent) *Command {
	if event.Button != 1 || event.MouseButtonEvent.State != sdl.RELEASED {
		return nil
	}

	mousePoint := sdl.Point{
		X: event.MouseButtonEvent.X,
		Y: event.MouseButtonEvent.Y,
	}

	return &Command{
		TypeOf:   ControlPointRelease,
		Target:   mousePoint,
		TargetId: cbc.index,
		Layer:    cbc.layer,
	}
}

func (cbc *CasteljauBezierCurve) cntlPointPressUpdate(cmd *Command) {
	if cbc.focused {
		return
	}
	cbc.focused = true
	cbc.index = cmd.TargetId

	cbc.ctlPoints[cbc.index].X, cbc.ctlPoints[cbc.index].Y = cmd.Target.X, cmd.Target.Y
	fmt.Printf("PRESSED x: %d, y %d\n", cbc.ctlPoints[cbc.index].X, cbc.ctlPoints[cbc.index].Y)
}

func (cbc *CasteljauBezierCurve) cntlPointRelease(cmd *Command) {
	if cmd.TargetId != cbc.index {
		fmt.Printf("NOT RELEASED x: %+v, index %d, %d, y %d\n", cmd, cbc.index, cbc.ctlPoints[cbc.index].X, cbc.ctlPoints[cbc.index].Y)

		return
	}
	fmt.Printf("RELEASED x: %+v, index %d, %d, y %d\n", cmd, cbc.index, cbc.ctlPoints[cbc.index].X,
		cbc.ctlPoints[cbc.index].Y)

	cbc.Draw()
	cbc.focused = false
	cbc.index = 0
}

func (cbc *CasteljauBezierCurve) mousePositionCommand(cmd *Command) {
	if cbc.focused {
		cbc.ctlPoints[cbc.index].X, cbc.ctlPoints[cbc.index].Y = cmd.Target.X, cmd.Target.Y
		cbc.Draw()
	}
}

func (cbc *CasteljauBezierCurve) canvasClickCommand(cmd *Command) {
	if cbc.focused {
		return
	}
	cbc.Add(cmd.Target)
	cbc.Draw()
}

func round(val float64) int32 {
	if val < 0 {
		return int32(val - 0.5)
	}
	return int32(val + 0.5)
}
