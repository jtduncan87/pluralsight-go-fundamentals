package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//ConcurrentFunc illustrates concurrency
func ConcurrentFunc() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)
	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Tide")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Roll")
	}()

	waitGrp.Wait() // wait until both go routines report back that they're done
}
