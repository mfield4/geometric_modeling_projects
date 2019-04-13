package ui

import (
	"sync"
)

var colMapInstance map[int]Ui
var once sync.Once

func GetCollisionMap() map[int]Ui {
	once.Do(func() {
		colMapInstance = map[int]Ui{}
	})
	return colMapInstance
}

var globalID = 0

func GUID() int {
	globalID++
	return globalID
}
