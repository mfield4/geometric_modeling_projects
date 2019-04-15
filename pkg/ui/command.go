package ui

type Event int

const (
	Noop Event = iota
	CanvasClick
	ControlPointPress
	ControlPointRelease
	MousePosition
)

type Mouse1Down interface {
	RegisterM1d(map[int]Mouse1Down)
	PressActive(x, y int32) bool
	Mouse1Down(x, y int32)
}

type Mouse1Up interface {
	RegisterM1u(map[int]Mouse1Up)
	ReleaseActive(x, y int32) bool
	Mouse1Up(x, y int32)
}

type MouseMotion interface {
	RegisterMM(map[int]MouseMotion)
	MotionActive() bool
	MouseMotion(x, y int32)
}
