package chapter02

import "fmt"

func myfunc1(arg ...int) {
	for _, v := range arg {
		fmt.Print(v, " ,")
	}
	fmt.Println()
}

func myfunc2(arg ...int) {
	myfunc1(arg...)
	myfunc1(arg[1:]...)
}

func myPrintf(arg ...interface{}) {
	for _, v := range arg {
		switch v.(type) {
		case int:
			fmt.Println(v, " is int")
		case int64:
			fmt.Println(v, " is int64")
		case string:
			fmt.Println(v, " is string")
		default:
			fmt.Println(v, " is unknow type")
		}
	}
}

func runvarg1() {
	myfunc2(1, 2, 3, 4, 5, 6)

	var v1 int = 3
	var v2 int64 = 55
	var v3 string = "123"
	var v4 float32 = 1.2
	myPrintf(v1, v2, v3, v4)
}

func Run() {
	// runvarg1()
	// runmultireturn()
	// runclosur()
	rundefer(1)
}
