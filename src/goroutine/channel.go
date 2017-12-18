package goroutine

import (
	"fmt"
	"time"
)

var ()

func Run3() {
	wg.Add(1)

	runner := make(chan int)
	go run(runner)

	runner <- 1
	wg.Wait()
}

func run(runner chan int) {
	p := <-runner
	if p > 4 {
		wg.Done()
		fmt.Println("over")
		return
	}

	fmt.Printf("第%d名正在拼命的跑\n", p)
	time.Sleep(1000 * time.Millisecond)
	go run(runner)

	runner <- (p + 1)

}
