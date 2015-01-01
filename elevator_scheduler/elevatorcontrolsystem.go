package main

import "fmt"
import "math"

type request struct {
	from    int
	goingUp bool
	waiting int
	tick    int
}
type elevatorcontrolsystem struct {
	elevators []elevator
	maxFloors int
	minFloors int
	requests  []request
	tick      int
}

func InitializeAndGet(elevators int, minFloors int, maxFloors int) *elevatorcontrolsystem {
	var e elevatorcontrolsystem
	for i := 0; i < elevators; i++ {
		e.elevators = append(e.elevators, *NewElevator(i))
	}
	e.minFloors = minFloors
	e.maxFloors = maxFloors
	return &e
}
func (e *elevatorcontrolsystem) Status() {
	for i := range e.elevators {
		fmt.Print("\nElevator# ", i, " Current Floor# ", e.elevators[i].currentfloor, " Goal Floors: ")
		for j := range e.elevators[i].goalfloor {
			fmt.Print(" ", j)
		}
	}
}
func (e *elevatorcontrolsystem) PickUp(fromWhere int, direction int) {
	var req request
	if direction == 1 {
		req.goingUp = true
	} else {
		req.goingUp = false
	}
	req.from = fromWhere
	req.waiting = 0
	req.tick = e.tick
	e.requests = append(e.requests, req)
}
func (e *elevatorcontrolsystem) Update(elevatorNo int, goalFloorNo int) {
	e.elevators[elevatorNo].goalfloor[goalFloorNo] = true
}

func (e *elevatorcontrolsystem) Step() {

	//This tick is used to keep the pickup request from starving
	e.tick++

	for i := range e.elevators {

		//let idle elevators be idle
		if len(e.elevators[i].goalfloor) > 0 {
			if e.elevators[i].goingUp == true {
				if e.elevators[i].currentfloor < e.maxFloors {
					e.elevators[i].currentfloor++
				} else {
					//The elevator hits the roof lets go down
					e.elevators[i].currentfloor--
					e.elevators[i].goingUp = false
				}

			} else {
				if e.elevators[i].currentfloor > e.minFloors {
					e.elevators[i].currentfloor--
				} else {
					//The elevator hits the basement, now lets glo up
					e.elevators[i].currentfloor++
					e.elevators[i].goingUp = true

				}
			}
		}

		//clear the workload from the current floor
		if e.elevators[i].goalfloor[e.elevators[i].currentfloor] == true {
			fmt.Println("Halting at ", e.elevators[i].currentfloor, " from elevator# ", i)

			delete(e.elevators[i].goalfloor, e.elevators[i].currentfloor)
		}

	}

	//Our elevator scheduler algo starts
	for j := range e.requests {
		var closestElevatorNumber int = -1
		var minDist float64 = 20000

		//First we try to find an elevator which is closest elevator going in the same direction as the request
		for i := range e.elevators {

			if e.requests[j].goingUp == true {
				if e.elevators[i].goingUp == e.requests[j].goingUp && e.elevators[i].currentfloor <= e.requests[j].from {
					dist := math.Abs(float64(e.requests[j].from - e.elevators[i].currentfloor))
					if dist < minDist {
						minDist = dist
						closestElevatorNumber = i
					}
				}
			} else {
				if e.elevators[i].goingUp == e.requests[j].goingUp && e.elevators[i].currentfloor >= e.requests[j].from {
					dist := math.Abs(float64(e.requests[j].from - e.elevators[i].currentfloor))
					if dist < minDist {
						minDist = dist
						closestElevatorNumber = i
					}
				}

			}
		}

		if closestElevatorNumber == -1 {
			fmt.Println("Assigning idle elevator")
			//we have not found closest elevator in the same direction
			//lets assign a closest idle elevator
			minDist = 20000
			for i := range e.elevators {
				if len(e.elevators[i].goalfloor) == 0 {
					dist := math.Abs(float64(e.requests[j].from - e.elevators[i].currentfloor))
					if dist < minDist {
						minDist = dist
						closestElevatorNumber = i
					}
				}
			}
		}

		if closestElevatorNumber == -1 {
			if e.tick-e.requests[j].tick >= 5 {
				//there was no idle elevator available
				//we wait till the request gets starved (ticks difference >5) and assign a elevator forcefully
				//To be implemented
			}
		}
		fmt.Println("Assigning elevator #", closestElevatorNumber, "to request ", j)

		if closestElevatorNumber != -1 {
			e.elevators[closestElevatorNumber].goalfloor[e.requests[j].from] = true
			e.requests = append(e.requests[:j], e.requests[j+1:]...)
			//todo delete the jth request
		}

	}
}
