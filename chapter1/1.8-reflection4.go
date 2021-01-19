package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Go Web Program"
	v1 := reflect.ValueOf(&name)
	fmt.Println("v1 可写性为:", v1.CanSet())

	v2 := v1.Elem()
	fmt.Println("v2 可写性为:", v2.CanSet())
}
