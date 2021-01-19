//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++
package main

import "fmt"

func main() {

	//str := "hello string!"
	//str = "hello modify!"
	//fmt.Println(str)
	//

	//	str1 := "\"Go Web\",I love you \n"//支持转义，但不能用来引用多行
	//	str2 :=`"Go Web",
	//I love you \n`  //支持多行组成，但不支持转义
	//	println(str1)
	//	println(str2)

	str := "I love" + " Go Web"
	str += " programming"
	fmt.Println(str) // I love Go Web programming

	//str := "programming"
	//fmt.Println(str[1]) //获取字符串索引位置为1的原始字节，比如r为114
	//fmt.Println(str[1:3]) //截取得字符串索引位置为 1 到 2 的字符串
	//fmt.Println(str[1:]) //截取得字符串索引位置为 1 到 len(s)-1 的字符串
	//fmt.Println(str[:3]) //截取得字符串索引位置为 0 到 2 的字符串
	//fmt.Println(len(str)) //获取字符串的字节数
	//fmt.Println(utf8.RuneCountInString(str)) //获取字符串字符的个数
	//fmt.Println([]rune(str)) // 将字符串的每一个字节转换为码点值，比如这里会输出[112 114 111 103 114 97 109 109 105 110 103]
	//fmt.Println(string(str[1])) // 获取字符串索引位置为1的字符值

	//var buffer bytes.Buffer  //创建一个空的 bytes.Buffer
	//for   {
	//	if piece,ok := getNextString();ok {
	//		buffer.WriteString(piece)  //通过 WriteString 方法将我们需要串联的字符串写入到 buffer 中
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println(buffer.String())  //最后用于取回整个级联的字符串

	//str := "love go web"
	//for index, char := range str {
	//	fmt.Printf("%d %U %c \n", index, char, char)
	//}

	//str := "go web"
	//fmt.Println(string(str[0])) //获取索引为0的字符

	//str := "i love go web"
	//str := `i love go web`
	//chars := []rune(str) //先把字符串转为rune切片
	//for _, char := range chars {
	//	fmt.Println(string(char))
	//}

	//str := "Hi 世界！"
	//by := []byte(str)    // 转换为 []byte，数据被自动复制
	//by[2] = ','        // 把空格改为半角逗号
	//fmt.Printf("%s\n", str)
	//fmt.Printf("%s\n", by)

	//str := "Hi 世界"
	//by := []rune(str)    // 转换为[]rune，数据被自动复制
	//by[3] = '中'
	//by[4] = '国'
	//fmt.Println(str)
	//fmt.Println(string(by))

}

func getNextString() (string, bool) {
	return "", true
}
