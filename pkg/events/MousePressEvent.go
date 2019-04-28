package events

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"sort"
)

type MousePressObserver interface {
	Layer(state, x, y int32) int
	Press(state, x, y int32)
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
	x, y, state := sdl.GetMouseState()

	sort.Slice(mpe.observers, func(i, j int) bool {
		return mpe.observers[i].Layer(int32(state), x, y) > mpe.observers[j].Layer(int32(state), x, y)
	})

	fmt.Printf("%+v\n", mpe.observers)

	if mpe.observers[0].Layer(int32(state), x, y) > 0 {
		mpe.observers[0].Press(int32(state), x, y)
	}
}

func (mpe *MousePressEvent) Subscribe(observer ...MousePressObserver) {
	for _, obs := range observer {
		mpe.observers = append(mpe.observers, obs)
	}
}
