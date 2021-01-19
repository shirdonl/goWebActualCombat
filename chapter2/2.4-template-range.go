package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//创建一个模版
	rangeTemplate := `
{{if .Kind}}
{{range $i, $v := .MapContent}}
{{$i}} => {{$v}} , {{$.OutsideContent}}
{{end}}
{{else}}
{{range .MapContent}}
{{.}} , {{$.OutsideContent}}
{{end}}    
{{end}}`

	str1 := []string{"第一次 range", "用 index 和 value"}
	str2 := []string{"第二次 range", "没有用 index 和 value"}

	type Content struct {
		MapContent     []string
		OutsideContent string
		Kind           bool
	}
	var contents = []Content{
		{str1, "第一次外面的内容", true},
		{str2, "第二次外面的内容", false},
	}

	// 创建模板并将字符解析进去
	t := template.Must(template.New("range").Parse(rangeTemplate))

	// 接收并执行模板
	for _, c := range contents {
		err := t.Execute(os.Stdout, c)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
