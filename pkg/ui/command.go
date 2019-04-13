package ui

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Event int

const (
	Noop Event = iota
	CanvasClick
)

type Command struct {
	typeOf   Event
	target   sdl.Point
	targetId int
}

func NewMouseButtonCommand(mouseEvent *sdl.MouseButtonEvent, targetId int) *Command {
	event := math.Abs(float64(mouseEvent.State-1)) + 1

	return &Command{
		typeOf: Event(event),
		target: sdl.Point{
			X: mouseEvent.X,
			Y: mouseEvent.Y,
		},
		targetId: targetId,
	}
}
