mesosphere
==========

Elevator Scheduling Algorithm

Algorithm:
This algorithm aspires to be better than FCFS. The basic logic is that the pickup request is assigned to the elevator which:
1. is going in the same direction
2. Closest among others in the same direction

If none such elevator is found, the closest idle elevator is assigned

If none such elevator is found, we wait for another cycle.

We maintain the time the request was added (as a tick counter) and the elevator systems tick counter is increased everytime the Step() is called.

This is to prevent starvation. When the pickup request is un-assigned for more than 5 cycles, the closest elevator is directly assigned to it.

This algorithm should be more optimal in case the number of floors is high e.g. 100

Instructions to run the program:

go to the main folder and run
go build

the executable name is elevator_scheduler

The basic usage is 
Usage: ./elevator_scheduler -n #Elevators -l lowest floor -u uppermost floor
elevator_scheduler 2 0 5

Then the commands mentioned in the problem statement can be issued:

status
pickup floor# direction
step
exit
printrequests

The elevator accepts multiple requests and it has been maintained in a map.

Unit test cases have not been written









