package main

import (
	"fmt"
	"math"
)

func main() {
	getPow := func(x float64, y float64) float64 { //声明函数变量
		return math.Pow(x, y)
	}
	fmt.Println(getPow(3, 4)) //使用函数
}
