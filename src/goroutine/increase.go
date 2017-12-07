package goroutine

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter int
	wg      sync.WaitGroup
)

func Run2() {
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go increase(1)
	go increase(2)
	wg.Wait()
	fmt.Println(counter)
}

func increase(id int) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}
