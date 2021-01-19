package main

import "fmt"

func main() {
	/* 定义局部变量 */
	num1 := 6
	num2 := 8
	fmt.Printf("交换前num1的值为 : %d\n", num1)
	fmt.Printf("交换前num2的值为 : %d\n", num2)
	/* 通过调用函数来交换值 */
	exchange(&num1, &num2)
	fmt.Printf("交换后num1的值 : %d\n", num1)
	fmt.Printf("交换后num2的值 : %d\n", num2)
}

/* 定义相互交换值的函数 */
func exchange(x *int, y *int) int {
	var tmp int
	tmp = *x /* 将 *x 值赋给 tmp */
	*x = *y  /* 将 *y 值赋给 *x */
	*y = tmp /* 将 tmp 值赋给 y*/
	return tmp
}
