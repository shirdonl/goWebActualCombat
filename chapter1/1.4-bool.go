package main

import "fmt"

func main() {
	var aVar = 100

	fmt.Println(aVar == 50)  //false
	fmt.Println(aVar == 100) //true
	fmt.Println(aVar != 50)  //true
	fmt.Println(aVar != 100) //false
	var str1 string
	var str2 string
	if str1 == "java" && str2 == "golang" {

	}

	//var c int
	//if 1 <= c && c <= 9 ||
	//	10 <= c && c <= 19 ||
	//	20 <= c && c <= 30 {
	//	//...
	//}

	i := 0
	b := true
	if b {
		i = 1
	}
	fmt.Println(i)
	fmt.Println(intToBool(0))

	var d bool
	fmt.Println(int(d) * 5)

}

//将bool转换为int
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func intToBool(i int) bool { return i != 0 }
