package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type PreviewCanvas struct {
	*Canvas
	pdst    *sdl.Rect
	current *CasteljauBezierCurve
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

	xScale := pc.dst.W / pc.pdst.W
	yScale := pc.dst.H / pc.pdst.H

	curCurve := pc.curves[pc.currentCurve].current()
	newCurve := make([]sdl.Point, len(curCurve))

	ox, oy := pc.pdst.X, pc.pdst.Y

	for i, pt := range curCurve {
		pt.X, pt.Y = ox+pt.X/xScale, oy+pt.Y/yScale
		newCurve[i] = pt
	}

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLines(newCurve)
}
