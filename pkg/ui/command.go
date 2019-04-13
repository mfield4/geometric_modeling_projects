package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Event int

const (
	Noop Event = iota
	CanvasClick
	ControlPointPress
	ControlPointRelease
	MousePosition
)

type Command struct {
	TypeOf   Event
	Target   sdl.Point
	TargetId int
	Layer    int
}

type MouseEvent struct {
	*sdl.MouseButtonEvent
	*sdl.MouseMotionEvent
}
