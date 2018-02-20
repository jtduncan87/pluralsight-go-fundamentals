package interfaces

import (
	"fmt"
	"strconv"
)

// there are methods already scattered through this project
// examples can be found in functions.salutation.greet()
// incrementNumberOfDoors below illustrates a function receiver method

type house struct {
	numDoors   int
	numWindows int
}

func (home *house) incrementNumberOfDoors() {
	home.numDoors++
}
func (home house) getDoorText() string {
	doors := "door"
	if home.numDoors != 1 {
		doors = doors + "s"
	}
	return "The house has " + strconv.Itoa(home.numDoors) + " " + doors
}

//PointerReceiverMethodExample demonstrates defining a method on a pointer receiver
func PointerReceiverMethodExample() {
	home := house{1, 2}
	fmt.Println(home.getDoorText())
	for x := 0; x < 4; x++ {
		home.incrementNumberOfDoors()
		fmt.Println(home.getDoorText())
	}
}
