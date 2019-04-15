package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Menu struct {
	preview *PreviewCanvas
}

func NewMenu(width int32, heigth int32, canvas *Canvas) *Menu {
	return &Menu{
		preview: NewPreviewCanvas(WindowWidth-width, heigth/2, width, heigth/2, canvas),
	}
}

func (m *Menu) Render(renderer *sdl.Renderer) {
	m.preview.Render(renderer)
}

func (*Menu) Rect() *sdl.Rect {
	return nil
}

func (*Menu) RegisterCol(colM map[int]Ui) {
}

func (menu *Menu) RegisterM1d(downs map[int]Mouse1Down) {

}
