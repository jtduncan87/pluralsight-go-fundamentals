package structs

import "fmt"

type familyMember struct {
	FirstName string
	LastName  string
	Age       int
}

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
