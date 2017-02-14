package main

import (
	"chapter02ex"
	"fmt"
	// "mymath"
	"flag"
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
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", *infile)
	}

	values, err := chapter02ex.ReadValues(*infile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(values)
		chapter02ex.WriteValues(values, *outfile)
	}
}
