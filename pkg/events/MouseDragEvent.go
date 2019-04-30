package events

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"sort"
)

type MouseDragObserver interface {
	Layer(state, x, y int32) int
	Drag(x, y int32)
}

type MouseDragEvent struct {
	observers []MouseDragObserver
}

func NewMouseDragEvent() *MouseDragEvent {
	return &MouseDragEvent{
		observers: nil,
	}
}

func (mpe *MouseDragEvent) Notify() {
	x, y, state := sdl.GetMouseState()

	sort.Slice(mpe.observers, func(i, j int) bool {
		return mpe.observers[i].Layer(int32(state), x, y) > mpe.observers[j].Layer(int32(state), x, y)
	})

	fmt.Printf("%+v\n", mpe.observers)

	if mpe.observers[0].Layer(int32(state), x, y) > 0 {
		mpe.observers[0].Drag(x, y)
	}
}

func (mpe *MouseDragEvent) Subscribe(observer ...MouseDragObserver) {
	for _, obs := range observer {
		mpe.observers = append(mpe.observers, obs)
	}
}
