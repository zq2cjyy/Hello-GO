package chapter03

import "fmt"

type Rect struct {
	x, y          float64
	height, width float64
}

func (rect *Rect) Area() float64 {
	return rect.height * rect.width
}

func CreateRect(x, y, h, w float64) *Rect {
	return &Rect{x, y, h, w}
}

func Run() {
	rect1 := &Rect{0, 0, 100, 100}
	fmt.Println(rect1.Area())
}
