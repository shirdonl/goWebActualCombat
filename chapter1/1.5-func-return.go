package main

func change(a, b int) (x, y int) {
	x = a + 100
	y = b + 100

	return   //101, 102
	//return x, y  //同上
	//return y, x  //102, 101
}

func main(){
	a := 1
	b := 2
	c, d := change(a, b)
	println(c, d)
}