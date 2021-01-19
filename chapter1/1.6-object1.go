package main

import (
	"fmt"
)

// 三角形结构体
type Triangle struct {
	Bottom float32
	Height float32
}

// 计算三角形面积
func (t *Triangle) Area() float32 {
	return (t.Bottom * t.Height) / 2
}

func main() {
	r := Triangle{6, 8}
	// 调用 Area() 方法，计算面积
	fmt.Println(r.Area())
}
