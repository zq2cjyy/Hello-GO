package main

import (
	//"chapter02ex"
	"fmt"
	// "mymath"
	//"chapter02ex/algorithms/bubblesort"
	"flag"
	"chapter03"
	"chapter02"
	"interfaceex"
)

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

// func main() {

// 	// var a int = 15
// 	// var b int = 16
// 	// c := mymath.Add(a, b)
// 	// fmt.Println(c)
// 	chapter02.Run()
// }
func main() {
	//flag.Parse()
	//if infile != nil {
	//	fmt.Println("infile=", *infile)
	//}
	//
	//values, err := chapter02ex.ReadValues(*infile)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(values)
	//	bubblesort.BubbleSort(values)
	//	fmt.Println(values)
	//	chapter02ex.WriteValues(values, *outfile)
	//}
	//var arrayA = []int{1, 2, 3};
	//var arrayB = arrayA;
	//arrayB[1]++;
	//fmt.Println(arrayA);
	//rect := &chapter03.Rect{0,0,100,100}
	//fmt.Println(rect.Area())

	fmt.Println("Fuck")
	var rect = chapter03.CreateRect(0, 0, 100, 101)
	fmt.Println(rect.Area())
	chapter02.Run();
	interfaceex.Run()
}
