//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
	"regexp"
)

func main()  {
	re := regexp.MustCompile(`who(o*)a(a|m)i`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-whooooaai-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-whoami-"))

	//
	//re := regexp.MustCompile(`w(a*)i`)
	//fmt.Printf("%q\n", re.FindAllStringSubmatch("-wi-", -1))
	//fmt.Printf("%q\n", re.FindAllStringSubmatch("-waaai-", -1))
	//fmt.Printf("%q\n", re.FindAllStringSubmatch("-wi-wai-", -1))
	//fmt.Printf("%q\n", re.FindAllStringSubmatch("-waaaaai-wi-", -1))

	//re := regexp.MustCompile(`wh(|o)`)
	//fmt.Println(re.FindString("who"))
	//re.Longest()
	//fmt.Println(re.FindString("who"))

	//re := regexp.MustCompile(`w(a*)i`)
	//re.Longest()
	//fmt.Println(re.FindString("-waaaaai-wi-")) // waaaaai，不会存在第一个w和最后一个i组合的过程。

	//re1 := regexp.MustCompile(`w.`)
	//fmt.Printf("%d\n", re1.NumSubexp())
	//
	//re2 := regexp.MustCompile(`(.*)((w)eb)(.*)w`)
	//fmt.Println(re2.NumSubexp())


	//re := regexp.MustCompile(`w(a*)i`)
	//fmt.Printf("%s\n", re.ReplaceAll([]byte("-wi-waaaaai-"), []byte("T")))
	//// $1表示匹配的第一个子串，这是wi的中间无字符串，所以$1为空，然后使用空去替换满足正则表达式的部分。
	//fmt.Printf("%s\n", re.ReplaceAll([]byte("-wi-waaaaai-"), []byte("$1")))
	//// "$1W"等价与"$(1W)"，值为空，将满足条件的部分完全替换为空。
	//fmt.Printf("%s\n", re.ReplaceAll([]byte("-wi-waaaaai-"), []byte("$1W")))
	//// ${1}匹配(x*)，保留
	//fmt.Printf("%s\n", re.ReplaceAll([]byte("-wi-waaaaai-"), []byte("${1}W")))



	//s := "Hello Gopher, let's go!"
	////定义一个正则表达式reg，匹配Hello或者Go
	//reg := regexp.MustCompile(`(Hell|G)o`)
	//
	//s2 := "2020-08-01,this is a test"
	////定义一个正则表达式reg2,匹配 YYYY-MM-DD 的日期格式
	//reg2 := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	//
	////最简单的情况，用“T替换”"-ab-axxb-"中符合正则"a(x*)b"的部分
	//reg3 := regexp.MustCompile("w(a*)i")
	//fmt.Println(reg3.ReplaceAllString("-wi-waai-", "T"))
	//
	////${1}匹配"2020-08-01,this is a test"中符合正则`(Hell|G)`的部分并保留，去掉"Hello"与"Go"中的'o'并用"hi"追加
	//rep1 := "${1}hi"
	//fmt.Printf("%q\n", reg.ReplaceAllString(s, rep1))
	//
	////首先，"2020-08-01,this is a test"中符合正则表达式`(\d{4})-(\d{2})-(\d{2})`的部分是"2020-08-01",将该部分匹配'(\d{4})'的'2020'保留，去掉剩余部分
	//rep2 := "${1}"
	//fmt.Printf("%q\n", reg2.ReplaceAllString(s2,rep2))
	//
	////首先，"2020-08-01,this is a test"中符合正则表达式`(\d{4})-(\d{2})-(\d{2})`的部分是"2020-08-01",将该部分匹配'(\d{2})'的'08'保留，去掉剩余部分
	//rep3 := "${2}"
	//fmt.Printf("%q\n", reg2.ReplaceAllString(s2,rep3))
	//
	////首先，"2020-08-01,this is a test"中符合正则表达式`(\d{4})-(\d{2})-(\d{2})`的部分是"2020-08-01",将该部分匹配'(\d{2})'的'01'保留，去掉剩余部分,并追加"13:30:12"
	//rep4 := "${3}:15:25:12"
	//fmt.Printf("%q\n", reg2.ReplaceAllString(s2,rep4))

}
