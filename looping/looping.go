package looping

import (
	"fmt"
	"time"
)

// Timer illustrates looping
func Timer() {
	for _timer := 10; _timer >= 0; _timer-- {
		if _timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(_timer)
		time.Sleep(1 * time.Millisecond)
	}
}

//TimerWithLabel illustrates break to label in loop
func TimerWithLabel() {
outerloop:
	for i := 0; i <= 3; i++ {
		for _timer := 10; _timer >= 0; _timer-- {
			if i == 2 {
				break outerloop
			} else if _timer == 0 {
				fmt.Println("Boom!")
				break
			}
			fmt.Println(_timer)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

//RangeExample illustrates a for in range loop
func RangeExample() {
	booksOnShelf := []string{"Desiring God", "Satan Cast Out", "The Prayer that Changed the World", "Holy Bible"}
	booksRead := []string{"Desiring God", "Holy Bible"}
	for _, book := range booksOnShelf {
		for _, completed := range booksRead {
			if book == completed {
				fmt.Println("I have", book, "on the shelf, and I've read that one!")
			}
		}
	}
}
