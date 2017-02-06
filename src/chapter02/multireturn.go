package chapter02

import (
	"fmt"
)

func func1() (a int, b string) {
	return
}

func func2() (a int, b string) {
	a = 100
	b = "2gui"
	return
}

func func3() (a int, b string) {
	return 89, "xgui"
}

func runmultireturn() {
	a, b := func1()
	fmt.Println(a)
	fmt.Println(b)

	a, b = func2()
	fmt.Println(a)
	fmt.Println(b)

	a, b = func3()
	fmt.Println(a)
	fmt.Println(b)
}
