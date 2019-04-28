package events

import "github.com/veandco/go-sdl2/sdl"

type MousePressObserver interface {
	MousePressed(x, y int32)
}

type MousePressEvent struct {
	observers []MousePressObserver
}

func NewMousePressEvent() *MousePressEvent {
	return &MousePressEvent{
		observers: nil,
	}
}

func (mpe *MousePressEvent) Notify() {
	x, y, _ := sdl.GetMouseState()

	for _, obs := range mpe.observers {
		obs.MousePressed(x, y)
	}
}

func (mpe *MousePressEvent) Subscribe(observer ...MousePressObserver) {
	for _, obs := range observer {
		mpe.observers = append(mpe.observers, obs)
	}
}
