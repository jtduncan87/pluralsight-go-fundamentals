package concurrency

import (
	"fmt"
)

// ChannelExample demonstrates channels
func ChannelExample() {
	done := make(chan bool)
	var factorial float64 = 1
	go func() {
		for i := 1; i < 4; i++ {
			factorial = factorial * float64(i)
		}
		done <- true
	}() // anonymous closured function

	<-done // blocks until done... this can also be assigned to
	fmt.Printf("%f", float64(factorial))
}

//BufferedChannelExample demonstrates buffered channels
func BufferedChannelExample() {
	done := make(chan bool, 2)
	go func() {
		fmt.Println("Hey")
		done <- true
		done <- true
	}()
	<-done
	<-done
	fmt.Println("Since both dones have processed, bye!")
}

type cheers []string

func (chs cheers) cheerer(c chan string) {
	for _, s := range chs {
		c <- s // send to channel
	}
	close(c)
}

//IterateChannelExample does what is says
func IterateChannelExample() {
	_cheers := cheers{
		"Roll Tide",
		"Go Bama",
		"Rammer Jammer",
		"RTR!",
	}
	c := make(chan string)
	c2 := make(chan string)
	//Call a go routine to fill the channel
	go _cheers.cheerer(c)
	go _cheers.cheerer(c2)
	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")
		}
	}
}
