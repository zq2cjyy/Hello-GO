package main

import (
	"fmt"
)

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(),array value:", array)

}
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println(array)
	fmt.Println("input over!")
}
