package concurrentProgrammingInGo

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// GoRoutineMain demonstrates go keyword. Main function is itself a go routine
// so it exits before the child go routines can finish, unless we yield to them
func GoRoutineMain() {
	runtime.GOMAXPROCS(2)
	godur, _ := time.ParseDuration("10ms")
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Roll" + strconv.Itoa(i))
			time.Sleep(godur)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Tide!" + strconv.Itoa(i))
			time.Sleep(godur)
		}
	}()
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
