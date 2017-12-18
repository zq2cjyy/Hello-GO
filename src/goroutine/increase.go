package goroutine

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
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
	for i := 0; i < 100; i++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
		//atomic.AddInt64(&counter, 1)
		//runtime.Gosched()
	}
}
