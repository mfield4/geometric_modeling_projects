package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type PreviewCanvas struct {
	*Canvas
	pdst *sdl.Rect
}

func NewPreviewCanvas(x, y, width, height int32, canvas *Canvas) *PreviewCanvas {
	dst := &sdl.Rect{
		X: x,
		Y: y,
		W: width,
		H: height,
	}
	return &PreviewCanvas{
		Canvas: canvas,
		pdst:   dst,
	}
}

func (pc PreviewCanvas) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(10, 10, 10, 255)
	renderer.FillRect(pc.pdst)

	curCurve := pc.curves[pc.currentCurve]
	if curCurve == nil || len(curCurve.ctlPoints) == 0 {
		return
	}

	newCurve := make([]sdl.Point, len(curCurve.curvePoints))

	xScale := pc.dst.W / pc.pdst.W
	yScale := pc.dst.H / pc.pdst.H

	ox, oy := pc.pdst.X, pc.pdst.Y

	for i, pt := range curCurve.curvePoints {
		pt.X, pt.Y = ox+pt.X/xScale, oy+pt.Y/yScale
		newCurve[i] = pt
	}
	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLines(newCurve)

	l, r := curCurve.splitCurve(0.5, false)

	for i, pt := range l.ctlPoints {
		pt.X, pt.Y = ox+pt.X/xScale, oy+pt.Y/yScale
		l.ctlPoints[i] = pt
	}

	for i, pt := range r.ctlPoints {
		pt.X, pt.Y = ox+pt.X/xScale, oy+pt.Y/yScale
		r.ctlPoints[i] = pt
	}

	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawLines(l.ctlPoints)

	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawLines(r.ctlPoints)

}
