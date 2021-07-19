package wtt

import (
	"fmt"

	"github.com/fatih/color"
)

func Wtt() {
	color.Red("see you, this package is wtt for test go mod")
}

var Name string = "tom"

func init() {
	fmt.Println("wtt 准备就绪")
}
