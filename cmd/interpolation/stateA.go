package main

type StateA struct {
}

func NewStateA() *StateA {
	return &StateA{}
}

func (*StateA) Input() bool {
	panic("implement me")
}

func (*StateA) Collision() {
	panic("implement me")
}

func (*StateA) Update() {
	panic("implement me")
}

func (*StateA) Render() {
	panic("implement me")
}
