package main

func main() {

}

//获取整型数组中的最小值
func min(arr []int) (min int) {
	min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}
