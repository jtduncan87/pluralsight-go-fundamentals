package concurrentProgrammingInGo

func primeGenerate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func primeFilter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
