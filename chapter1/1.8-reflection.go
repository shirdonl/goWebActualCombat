package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var x float64 = 3.4
	//fmt.Println("type:", reflect.TypeOf(x))

	//var x float64 = 6.8
	//fmt.Println("value:", reflect.ValueOf(x)) //Valueof方法会返回一个Value类型的对象

	//var x float64 = 6.8
	//v := reflect.ValueOf(x)
	//fmt.Println("type:", v.Type())
	//fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//fmt.Println("value:", v.Float())

	//var x uint8 = 'x'
	//v := reflect.ValueOf(x)
	//fmt.Println("type:", v.Type())                            // uint8.
	//fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	//x = uint8(v.Uint())// v.Uint returns a uint64. 这个地方必须进行强制类型转换！

	type DefinedInt int
	var x DefinedInt = 8
	v := reflect.ValueOf(x)
	fmt.Println(v)
}
