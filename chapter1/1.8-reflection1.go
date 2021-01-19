package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Go Web Program"

	v := reflect.ValueOf(name)
	fmt.Println("可写性为:", v.CanSet())
}
