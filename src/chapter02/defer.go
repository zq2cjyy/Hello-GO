package chapter02

import (
	"fmt"
)

// defer 关键字 保证了defer后的方法一定会执行，执行实际在方法执行完即将退出之前，或者抛出异常即将退出之前
func rundefer(i int) {

	if i < 1 {
		return
	}
	defer func() {
		fmt.Println("i<1")
	}()

	if i > 1 {
		return
	}
	defer func() {
		fmt.Println("i>1")
	}()

	fmt.Println("i=1")
}
