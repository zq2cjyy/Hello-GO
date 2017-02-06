package chapter02

import (
	"fmt"
)

func runclosur() {
	a := func() func(mz string) {
		return func(m string) {
			fmt.Println("w2 ", m)
		}
	}()

	a("gui")
	a("cj")
}
