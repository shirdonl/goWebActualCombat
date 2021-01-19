package main

import "fmt"

func exchange2(c, d *int) {
	d, c = c, d
}
func main() {
	x, y := 6, 8
	exchange2(&x, &y)
	fmt.Println(x, y)
}
