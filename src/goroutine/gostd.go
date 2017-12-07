package goroutine

import (
	"sync"
	"fmt"
	"runtime"
)

func Run() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		//for c := 'a'; c <= 'a'+26; c++ {
		//	fmt.Printf("%c", c)
		//}
		for i := 0; i < 10000; i++ {
			fmt.Print("1")
		}
		fmt.Println()
	}()
	go func() {
		defer wg.Done()
		//for c := 'A'; c <= 'A'+26; c++ {
		//	fmt.Printf("%c", c)
		//}
		for i := 0; i < 10000; i++ {
			fmt.Print("2")
		}
		fmt.Println()
	}()
	wg.Wait()

	fmt.Println("DONE!")
}
