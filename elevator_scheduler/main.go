// elevator_scheduler project main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFromStdin() string {
	fmt.Print("# ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = line[0 : len(line)-1]
	return line
}

var flagNumElev = flag.Int("n", 0, "intflag")
var flagMinFloor = flag.Int("l", 0, "intflag")
var flagMaxFloor = flag.Int("u", 0, "intflag")

func main() {
	flag.Parse()
	if *flagNumElev <= 0 {
		fmt.Printf("Usage: %s -n #Elevators -l lowest floor -u uppermost floor\n", os.Args[0])
		os.Exit(1)
	}

	elevatorControlSystem := InitializeAndGet(*flagNumElev, *flagMinFloor, *flagMaxFloor)
	for {
		line := readFromStdin()

		if strings.Contains(line, "status") {
			fmt.Println("status")
			elevatorControlSystem.Status()

		} else if strings.Contains(line, "pickup") {

			args := strings.Split(line, " ")
			from, _ := strconv.ParseInt(args[1], 10, 64)
			dir, _ := strconv.ParseInt(args[2], 10, 64)
			fmt.Println("pickup ", from, dir)
			elevatorControlSystem.PickUp(int(from), int(dir))

		} else if strings.Contains(line, "step") {
			fmt.Println("step")
			elevatorControlSystem.Step()
		} else if strings.Contains(line, "exit") {
			fmt.Println("exit")
			return
		} else if strings.Contains(line, "printrequests") {
			fmt.Println("printrequests")
			fmt.Println(elevatorControlSystem.requests)
		} else {
			fmt.Println("invalid cmd")
		}
	}

}
