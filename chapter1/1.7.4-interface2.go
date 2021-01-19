package main

type stringer interface {
	string() string
}

type tester interface {
	stringer // 嵌入其他接口
	test()
}

type data struct{}

func (*data) test() {
}

func (data) string() string {
	return ""
}

func main() {
	var d data

	var t tester = &d
	t.test()
	println(t.string())
}
