package main

import (
	"fmt"
	"github.com/mfield4/178_projects/pkg/app"
	"github.com/mfield4/178_projects/pkg/curves"
	"github.com/mfield4/178_projects/pkg/events"
	"github.com/mfield4/178_projects/pkg/ui"
	"github.com/mfield4/178_projects/pkg/ui/canvas"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type StateA struct {
	MousePress *events.MousePressEvent
	MouseDrag  *events.MouseDragEvent
	ui         []ui.Ui
}

func NewStateA() *StateA {
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
	newState := &StateA{
		MousePress: events.NewMousePressEvent(),
		MouseDrag:  events.NewMouseDragEvent(),
	}
	mainCan := newState.NewMainCanvas()
	newState.NewRaisedCanvas(mainCan)
	newState.NewReducedCanvas(mainCan)
	newState.NewButton().NewButton().NewButton()

	return newState
}

/*
TODO
Each state will be a factory for all it's ui elements. Do this refactor in the following steps:
	1. Delete/Move all current constructors
	2. Refactor constructors to this file
	3. Profit??
*/
func (s *StateA) NewBezierCurve() *curves.CasteljauBezierCurve {
	cur := curves.NewBezierCurve(0, 2, true)
	s.MousePress.Subscribe(cur)
	s.MouseDrag.Subscribe(cur)
	return cur
}

func (s *StateA) NewMainCanvas() *canvas.MainCanvas {
	can := canvas.NewMainCanvas(0, 0, app.WindowWidth, app.WindowHeight/2, s.NewBezierCurve())

	s.MousePress.Subscribe(can)
	s.ui = append(s.ui, can)
	return can
}

func (s *StateA) NewRaisedCanvas(main canvas.Canvas) *canvas.RaisedCanvas {
	can := canvas.NewRaisedCanvas(main)
	s.ui = append(s.ui, can)
	return can
}

func (s *StateA) NewReducedCanvas(main canvas.Canvas) *canvas.ReducedCanvas {
	can := canvas.NewReducedCanvas(main)
	s.ui = append(s.ui, can)
	return can
}

func (s *StateA) NewButton( /*TODO Params*/ ) *StateA {
	return s
}

func (s *StateA) Input() (updates []func(), ok bool) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return updates, false
		case *sdl.MouseMotionEvent:

			updates = append(updates, func() { s.MouseDrag.Notify() })
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
