package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//ConcurrentFunc illustrates concurrency
func ConcurrentFunc() {
	runtime.GOMAXPROCS(2) // parallelism
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
