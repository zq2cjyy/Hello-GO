package main

import (
	"fmt"
	"mymath"
)

func main() {

	var a int = 15
	var b int = 16
	c := mymath.Add(a, b)
	fmt.Println(c)
}