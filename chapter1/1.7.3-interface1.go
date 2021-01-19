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

type InterfaceFile interface {
	Read()
	Write()
}

type InterfaceReader interface {
	Read()
}

type File struct {
}

func (f *File) Read() {

}

func (f *File) Write() {

}

func main() {
	f := new(File)

	var f1 InterfaceFile = f
	var f2 InterfaceReader = f1

	var f3 InterfaceReader = new(File)

	// 接口查询
	if f5, ok := f3.(InterfaceFile); ok {
		fmt.Println(f5)
	}

	// 询问接口它指向的对象是否是某个类型
	if f6, ok := f3.(*File); ok {
		fmt.Println(f6)
	}

	fmt.Println(f1, f2, f3)
}
