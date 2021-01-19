package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Go Web Program"
	fmt.Println("真实 name 的原始值为：", name)

	v1 := reflect.ValueOf(&name)
	v2 := v1.Elem()

	v2.SetString("Go Web Program2")
	fmt.Println("通过反射对象进行更新后，真实 name 变为：", name)
}
