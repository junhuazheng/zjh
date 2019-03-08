package main

import "fmt"

func main() {
	//1、声明变量
	var a bool
	fmt.Println("a0 = ", a)

	a = true
	fmt.Println("a = ", a)

	//2、自动推到类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)
}
