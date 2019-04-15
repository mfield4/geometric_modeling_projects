package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Menu struct{}

func NewMenu(width int32, heigth int32) *Menu {
	return &Menu{}
}

func (*Menu) Render(*sdl.Renderer) {}

func (*Menu) Rect() *sdl.Rect {
	return nil
}

func (*Menu) Register(colM map[int]Ui) {
}

func (menu *Menu) RegisterM1d(downs map[int]Mouse1Down) {

}
