package maps

import (
	"fmt"
)

//MapExample demonstrates maps
func MapExample() {
	religions := make(map[string]bool)
	religions["Christianity"] = false
	religions["Islam"] = false
	religions["Mormonism"] = false

	for key, value := range religions {
		fmt.Printf("\n%v is %v", key, value)
	}
	fmt.Println("\nOops, one of those values is false!")
	religions["Christianity"] = true
	for key, value := range religions {
		fmt.Printf("\n%v is %v", key, value)
	}
	fmt.Println("\nFixed it!")

}
