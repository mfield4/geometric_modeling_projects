package main

import (
	"fmt"
	"github.com/mfield4/178_projects/pkg/app"
	"github.com/mfield4/178_projects/pkg/events"
	"github.com/mfield4/178_projects/pkg/ui"
	"github.com/mfield4/178_projects/pkg/ui/canvas"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"sync"
)

type StateA struct {
	MousePress *events.MousePressEvent
	MouseDrag  *events.MouseDragEvent
	ui         []ui.Ui
}

var gsa sync.Once
var stateA *StateA

func GetStateA() *StateA {
	gsa.Do(func() {
		stateA = newStateA()
	})

	return stateA
}

func newStateA() *StateA {
	// Init all elements in the state here
	// Register with the events.
	/*
		Elements to init:
			Main canvas
			Raised Preview
			Reduced Preview
			3 State buttons
			Optional: sliders?
	*/
	mousePress := events.NewMousePressEvent()
	mouseDrag := events.GetMouseDragEvent()

	// TODO fix ui code to be more generic here ?
	mainCanvas := canvas.NewMainCanvas(0, 0, app.WindowWidth, app.WindowHeight/2, mousePress)
	//raisedCanvas := canvas.NewRaisedCanvas(mainCanvas)
	//reducedCanvas := canvas.NewReducedCanvas(mainCanvas)
	// TODO BUTTONS
	//stateA := button.NewButton()
	//stateB := button.NewButton()
	//stateC := button.NewButton()

	return &StateA{
		MousePress: mousePress,
		MouseDrag:  mouseDrag,

		// Everything that gets drawn in this state
		ui: []ui.Ui{
			mainCanvas,
		},
	}
}

func (s *StateA) Input() (updates []func(), ok bool) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return updates, false
		case *sdl.MouseMotionEvent:
		case *sdl.MouseButtonEvent:
			fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
				t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)

			updates = append(updates, func() { s.MousePress.Notify() })
		case *sdl.KeyboardEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
		default:
		}
	}

	return updates, true
}

func (*StateA) Update(fns []func()) {
	for _, fn := range fns {
		fn()
	}
}

func (s *StateA) Render(renderer *sdl.Renderer) {
	err := renderer.Clear()
	err = renderer.SetDrawColor(0, 0, 0, 255)
	err = renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: app.WindowWidth,
		H: app.WindowHeight,
	})
	if err != nil {
		log.Println(err)
		return
	}

	for _, ui := range s.ui {
		ui.Render(renderer)
	}

	renderer.Present()
}
