package concurrentProgrammingInGo

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	//LOGFILE location
	LOGFILE = "./resources/log.txt"
)

//SimulatingMutex using a mutex obj
func SimulatingMutex() {
	mutex := new(sync.Mutex)
	runtime.GOMAXPROCS(4)
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			mutex.Lock()
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				mutex.Unlock()
			}()
		}
	}
	fmt.Scanln()
}

//SimulatingMutex2 Using a channel as a mutex... prob not as efficient as using the actual mutex
func SimulatingMutex2() {
	runtime.GOMAXPROCS(4)
	mutex := make(chan bool, 1)
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			mutex <- true
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-mutex
			}()
		}
	}
	fmt.Scanln()
}

//LoggingToAFile how to sync logging to a file
func LoggingToAFile() {
	runtime.GOMAXPROCS(4)
	f, _ := os.Create(LOGFILE)
	f.Close()
	logCh := make(chan string, 50)
	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile(LOGFILE, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()
	mutex := make(chan bool, 1)
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				<-mutex
			}()
		}
	}
	fmt.Scanln()
}

//EventExample simulates Events
func EventExample() {
	b := makeButton()
	handlerOne := make(chan string)
	handlerTwo := make(chan string)
	b.addEventListener("click", handlerOne)
	b.addEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler 1: ", msg)
		}
	}()
	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler 2: ", msg)
		}
	}()

	b.triggerEvent("click", "Button Clicked!")
	fmt.Println("------> Removing Handler 2.....")
	b.removeEventListener("click", handlerTwo)
	b.triggerEvent("click", "Button Clicked Again!")
	fmt.Scanln()
}

//CallbackExample simulates callbacks
func CallbackExample() {
	po := new(purchaseOrder)
	po.Value = 200.34

	ch := make(chan *purchaseOrder)

	go savePO(po, ch)

	newPO := <-ch

	fmt.Printf("PO: %v\n", newPO)
	fmt.Scanln()
}

//PromiseExample simulates promises
func PromiseExample() {
	po := new(purchaseOrder)
	po.Value = 10.43
	savePOWithPromise(po, true).then(func(o interface{}) error {
		po := o.(*purchaseOrder)
		fmt.Printf("PO: %v\n", po)
		return nil
	}, func(err error) {
		fmt.Println("Failed to save po: ", err.Error())
	})

	fmt.Scanln()
}

//PipeAndFilterExample simulates filtering primes
func PipeAndFilterExample() {
	ch := make(chan int)
	go primeGenerate(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go primeFilter(ch, ch1, prime)
		ch = ch1

		if prime > 100 {
			return
		}
	}
}

//ExtractTransformAndLoadExample demeonstrates the ETL pattern
//original implementation ran 20s
//2nd with goroutines and channels: 13.92s
//3rd adding sync.WaitGroup and spawning subgo-routines: 52ms
func ExtractTransformAndLoadExample() {
	start := time.Now()
	extractCh := make(chan *order)
	transCh := make(chan *order)
	doneCh := make(chan bool)
	go extractOrders(extractCh)
	go transformOrders(extractCh, transCh)
	go loadOrders(transCh, doneCh)
	<-doneCh
	fmt.Println(time.Since(start))
}
