package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name interface{} = "shirdon"

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", name, name)

	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	// 从反射对象到接口变量
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)
}
