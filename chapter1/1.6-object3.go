package main

import (
	"./person"
	"fmt"
)

func main() {
	s := new(person.Student)
	s.name = "shirdon"
	s.Age = 22
	fmt.Println(s.GetAge())

}
