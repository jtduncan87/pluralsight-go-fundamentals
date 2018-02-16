package slices

import (
	"fmt"
	"strconv"
)

//SliceItUp demonstrates slices
func SliceItUp() {
	myBooks := make([]string, 5, 10)
	fmt.Println("Capacity:", cap(myBooks), "length:", len(myBooks))
	for index := range myBooks {
		myBooks[index] = "Book " + strconv.Itoa(index+1)
	}
	for _, book := range myBooks {
		fmt.Println(book, "is in my slice")
	}

	sliceOfBooks := myBooks[2:4]
	for _, book := range sliceOfBooks {
		fmt.Println(book, "is in my slice of slices")
	}

	fmt.Println("-----Lets add to our slice of slices!-----")
	for i := 5; i < 16; i++ {
		sliceOfBooks = append(sliceOfBooks, "book "+strconv.Itoa(i))
		fmt.Printf("Capacity of the slice of slices is %d\n", cap(sliceOfBooks))
	}
}

//SliceReferenceExamples illustrates slice references and appends
func SliceReferenceExamples() {
	mySlice := []int{1, 3, 4}
	fmt.Println(mySlice)

	newSlice := []int{54, 20, 5}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

}
