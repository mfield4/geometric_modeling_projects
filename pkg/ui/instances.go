package ui

import (
	"sync"
)

type CommandMaps struct {
	m1dCommands     map[int]Mouse1Down
	m1uCommands     map[int]Mouse1Up
	m2dCommands     map[int]Mouse2Down
	mMotionCommands map[int]MouseMotion
}

var Commands = CommandMaps{}

var colMapInstance map[int]Ui

var col sync.Once

func GetCollisionMap() map[int]Ui {
	col.Do(func() {
		colMapInstance = map[int]Ui{}
	})
	return colMapInstance
}

var m1d sync.Once

func GetMouse1dCommands() map[int]Mouse1Down {
	m1d.Do(func() {
		Commands.m1dCommands = map[int]Mouse1Down{}
	})
	return Commands.m1dCommands
}

var m1u sync.Once

func GetMouse1uCommands() map[int]Mouse1Up {
	m1u.Do(func() {
		Commands.m1uCommands = map[int]Mouse1Up{}
	})
	return Commands.m1uCommands
}

var m2d sync.Once

func GetMouse2dCommands() map[int]Mouse2Down {
	m2d.Do(func() {
		Commands.m2dCommands = map[int]Mouse2Down{}
	})
	return Commands.m2dCommands
}

var mm sync.Once

func GetMouseMMCommands() map[int]MouseMotion {
	mm.Do(func() {
		Commands.mMotionCommands = map[int]MouseMotion{}
	})
	return Commands.mMotionCommands
}

var globalID = 0

func GUID() int {
	globalID++
	return globalID
}
