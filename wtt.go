package wtt

import "fmt"

func Wtt() {
	fmt.Println("see you, this package is wtt for test go mod")
}

var Name string = "tom"

func init() {
	fmt.Println("wtt 准备就绪")
}
