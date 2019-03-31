package main

import (
	"fmt"
	"time"
)

//打印机属于公共资源
//定义一个打印机，参数为字符串，按每个字符打印
func Printer(str string) {

	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
}

func person2() {
	Printer("world")
}

func main() {

	//新建两个协程，代表2个人，2个人同时使用打印机
	go person1()

	go person2()

	for {
	}
}
