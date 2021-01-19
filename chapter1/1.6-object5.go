//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import "fmt"

// 正方形
type Square struct {
	sideLen float32
}

// 三角形结构体
type Triangle struct {
	Bottom float32
	Height float32
}

// 计算三角形面积
func (t *Triangle) Area() float32 {
	return (t.Bottom * t.Height) / 2
}

// 接口 Shape
type Shape interface {
	Area() float32
}

// 计算正方形的面积
func (sq *Square) Area() float32 {
	return sq.sideLen * sq.sideLen
}

func main() {
	t := &Triangle{6, 8}
	s := &Square{8}
	shapes := []Shape{t, s}    // 创建一个 Shape 类型的数组
	for n, _ := range shapes { // 迭代数组上的每一个元素并调用 Area() 方法
		fmt.Println("图形数据: ", shapes[n])
		fmt.Println("它的面积是: ", shapes[n].Area())
	}
}
