package events

type Command struct {
	// info necessary for collision
	layer  int
	notify func()
}
