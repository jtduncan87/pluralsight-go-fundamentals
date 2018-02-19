package maps

import (
	"fmt"
)

var prefixMap = map[string]string{
	"Jacob":  "Mr",
	"Taylor": "Mrs",
	"Bob":    "Burger",
}

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

//MapInitShortand demonstrates shorthand map initialization
func MapInitShortand() {
	for key, value := range prefixMap {
		println(value, key)
	}
}

//DeleteFromMap demonstrates deleting from a map
func DeleteFromMap() {

	fmt.Println("before delete:")
	for key, value := range prefixMap {
		fmt.Println(value, key)
	}
	fmt.Println("After delete")
	delete(prefixMap, "Bob")
	for key, value := range prefixMap {
		fmt.Println(value, key)
	}
}

//CheckMapForExistance demonstrates checking for existence
func CheckMapForExistance() {
	bill := "bill"
	jake := "Jacob"

	doesItExist(prefixMap, bill)

	doesItExist(prefixMap, jake)
}

func doesItExist(someMap map[string]string, key string) {
	if _, exists := someMap[key]; exists {
		fmt.Println(key, "exists in the map")
	} else {
		fmt.Println(key, "doesn't exist in the map")
	}
}
