package events

import (
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

type MouseDragObserver interface {
	MouseDrag(x, y int32)
}

type MouseDragEvent struct {
	observers []MouseDragObserver
}

var mde sync.Once
var mouseDragEvent *MouseDragEvent

func GetMouseDragEvent() *MouseDragEvent {
	mde.Do(func() {
		mouseDragEvent = &MouseDragEvent{
			observers: nil,
		}
	})

	return mouseDragEvent
}

func (mpe *MouseDragEvent) Notify() {
	x, y, _ := sdl.GetMouseState()

	for _, obs := range mpe.observers {
		obs.MouseDrag(x, y)
	}
}

func (mpe *MouseDragEvent) Subscribe(observer ...MouseDragObserver) {
	for _, obs := range observer {
		mpe.observers = append(mpe.observers, obs)
	}
}
