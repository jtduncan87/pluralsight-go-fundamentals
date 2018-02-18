package functions

import (
	"fmt"
	"strings"
)

type salutation struct {
	name     string
	greeting string
}

//printer accepts a string
type printer func(string)

func (sal salutation) greet() (string, string) {
	return sal.greeting + " " + sal.name, "Hey! " + sal.name
}

func createPrintFunction(custom string) printer {
	return func(s string) { // closure
		fmt.Println(s + custom)
	}
}

//GreetMain demonstrates a function operating on a type
func GreetMain() {
	s := salutation{
		name:     "Jacob",
		greeting: "Roll Tide!",
	}
	greet1, greet2 := s.greet()
	fmt.Println("First return:", greet1)
	fmt.Println("Second return:", greet2)

	s2 := salutation{"Taylor", "Love ya!"}
	greetWithPassedFunc(s2, print)

	greetWithPassedFunc(s2, createPrintFunction("xoxo"))
}

func greetWithPassedFunc(sal salutation, do printer) {
	greet, alternate := sal.greet()
	do(greet)
	do(alternate)
}

func print(s string) {
	fmt.Println(s)
}

// Converter module to titleCase and author to Upper
func Converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)
	return module, author
}
