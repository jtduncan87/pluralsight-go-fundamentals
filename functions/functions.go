package functions

import (
	"fmt"
	"strings"
)

type salutation struct {
	name     string
	greeting string
}

func (sal salutation) greet() (string, string) {
	return sal.greeting + " " + sal.name, "Hey! " + sal.name
}

//GreetFunc demonstrates a function operating on a type
func GreetFunc() {
	s := salutation{
		name:     "Jacob",
		greeting: "Roll Tide!",
	}
	greet1, greet2 := s.greet()
	fmt.Println("First return:", greet1)
	fmt.Println("Second return:", greet2)
}

// Converter module to titleCase and author to Upper
func Converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)
	return module, author
}
