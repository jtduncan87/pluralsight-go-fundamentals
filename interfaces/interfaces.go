package interfaces

import (
	"fmt"
)

type sport interface {
	cheer() string
}

type baseball struct{}
type football struct{}
type basketball struct{}

func (b baseball) cheer() string {
	return "Root Root for the hometeam!"
}
func (b football) cheer() string {
	return "Roll Tide!"
}
func (b basketball) cheer() string {
	return "March Madness yo!"
}

func cheerForTheSport(isport sport) string {
	return isport.cheer()
}

//InterfaceExample demonstrates the use of interfaces
func InterfaceExample() {
	var sports [3]sport
	sports[0] = football{}
	sports[1] = baseball{}
	sports[2] = basketball{}

	for i := range sports {
		fmt.Println(cheerForTheSport(sports[i]))
	}
}
