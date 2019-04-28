package state

// Concrete states will consist of a collection of embedded events, then a collection of render components.
type State interface {
	Input() bool
	Collision()
	Update()
	Render()
}
