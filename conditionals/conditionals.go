package conditionals

import (
	"fmt"
	"os"
)

//SomeType for passing to SwitchOnType
type SomeType int

//IfElseAndSwitchExample illustrates control flow
func IfElseAndSwitchExample() {

	// if and else
	if firstCompare, secondCompare := 1, 3; firstCompare < (secondCompare - 1) {
		fmt.Println("First compare wins!")
	} else {
		fmt.Println("Second wins!")
	}

	// switch
	switch "roll" {
	case "roll", "tide":
		fmt.Println("tide")
		fallthrough
	case "derp":
		fmt.Println("roll")
	default:
		fmt.Println("chopon")
	}

	//error checking with if
	_, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("There was an error:", err)
	} else {
		fmt.Println("There wasn't error:", err)
	}
}

//SwitchOnType illustrates switching on types
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case SomeType:
		fmt.Println("SomeType")
	default:
		fmt.Println("Unknown")
	}
}
