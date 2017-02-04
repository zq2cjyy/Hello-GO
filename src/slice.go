package main

import "fmt"

func main() {

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println()
	var mySlice []int = myArray[:5]
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
