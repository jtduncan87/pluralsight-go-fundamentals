package pointers

import (
	"fmt"
	"os"
)

// PointerExample illustrates pointers
func PointerExample() {
	name := os.Getenv("LOGNAME")
	course := "Systematic Theology"
	fmt.Println("\nHi", name, "you're currently watching", course)
	changeCourse(&course)
	fmt.Println("\nYou're now watching course", course)
}

func changeCourse(course *string) string {
	*course = "Intro to Apologetics"

	fmt.Println("\nTrying to change your course to", *course)

	return *course
}
