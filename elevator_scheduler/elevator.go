package main

type elevator struct {
	id           int
	currentfloor int
	goingUp      bool
	goalfloor    map[int]bool
}

func NewElevator(ID int) *elevator {
	e := elevator{id: ID, goingUp: true}
	e.goalfloor = make(map[int]bool)
	e.currentfloor = -1
	return &e
}
