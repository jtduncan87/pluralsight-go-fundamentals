package slices

import (
	"fmt"
	"strconv"
)

type sliceWrapper []int

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

//SliceReferenceExamples illustrates slice references and appends, also demonstrates shorthand of slice initialization
func SliceReferenceExamples() {
	mySlice := []int{1, 3, 4}
	fmt.Println(mySlice)

	newSlice := []int{54, 20, 5}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

	fmt.Println("Changing index 1 of slice: ")
	newSlice[1] = 4
	fmt.Println(newSlice)

}

type sliceOfInts []int

//DeleteFromSliceExample illustrates deleting from a slice
func DeleteFromSliceExample() {
	slice := sliceOfInts{1, 2, 3, 4}
	fmt.Println("Before:", slice)
	//slice.deleteFromSlice(1)
	slice = append(slice[:1], slice[1+1:]...)
	fmt.Println("After:", slice)
}
