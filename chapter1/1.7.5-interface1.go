package main

func main() {
	var var1, var2 interface{}
	println(var1 == nil, var1 == var2)
	var1, var2 = 66, 88
	println(var1 == var2)
	var1, var2 = map[string]string{}, map[string]string{}
	println(var1 == var2)
}
