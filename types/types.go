package types

import (
	"fmt"
	"reflect"
)

const pi = 3.14

const (
	pi2 = 3.14
)

type familyMember struct {
	FirstName string
	LastName  string
	Age       int
}

//Name has string underlying it
type name string

//StructExample illustrates structs
func StructExample() {
	var famMember familyMember
	famMember.Age = 30
	famMember.FirstName = "Jake"
	famMember.LastName = "Dunc"
	fam := new(familyMember)
	fam.FirstName = "Bill"
	fam.LastName = "Cowpoke"
	fam2 := familyMember{
		Age:       30,
		FirstName: "Bob",
		LastName:  "Brosky",
	}
	fmt.Println(famMember)
	fmt.Println(fam)
	fmt.Println(fam2)
}

//TypeDemo prints TypeOf info of name:Name
func TypeDemo() {
	var _name name = "Jacob"
	fmt.Println(_name, "is of type", reflect.TypeOf(_name))
}

//ConstDemo demonstrates the usage of consts
func ConstDemo() {
	fmt.Println(pi, "is from a single const-prefixed declaration;", pi2, "is from a () declaration")
}
