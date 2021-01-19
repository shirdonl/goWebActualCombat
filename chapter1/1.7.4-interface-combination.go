package main

//接口1
type Interface1 interface {
	Write(p []byte) (n int, err error)
}

//接口2
type Interface2 interface {
	Close() error
}

//接口组合
type InterfaceCombine interface {
	Interface1
	Interface2
}
