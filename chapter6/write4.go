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
	"os"
)

func main() {
	//新建文件
	fout, err := os.Create("./write4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()
	for i := 0; i < 5; i++ {
		outstr := fmt.Sprintf("%s:%d\r\n", "Hello Go", i) //Sprintf格式化
		// 写入文件
		fout.WriteString(outstr)            //string信息
		fout.Write([]byte("i love go\r\n")) //byte类型
	}
}
