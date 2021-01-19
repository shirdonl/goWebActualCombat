package testexample

func Min(arr []int) (min int) { //获取整型数组中的最小值
	min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}
