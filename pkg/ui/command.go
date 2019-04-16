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
	RegisterM1d()
	PressActive(x, y int32) bool
	Mouse1Down(x, y int32)
}

type Mouse1Up interface {
	RegisterM1u()
	ReleaseActive(x, y int32) bool
	Mouse1Up(x, y int32)
}

type MouseMotion interface {
	RegisterMM()
	MotionActive() bool
	MouseMotion(x, y int32)
}

type Mouse2Down interface {
	RegisterM2d()
	RightActive() bool
	Mouse2Down(x, y int32)
}
