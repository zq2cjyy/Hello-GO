package mystudy

import (
	"fmt"
)

func main() {

	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println()
	var mySlice []int = myArray[:5]

	mySlice = append(mySlice, 6, 6, 6)

	mySlice2 := []int{8, 8, 8}
	//将myslice2复制到myslice结尾
	mySlice = append(mySlice, mySlice2...)
	//复制内容
	copy(mySlice, mySlice2)
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	//输出 数组切片的长度 与容量
	fmt.Println()
	fmt.Println("cap:", cap(mySlice))
	fmt.Println("len:", len(mySlice))
	fmt.Println()
}
