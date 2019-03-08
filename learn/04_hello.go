package main

import "fmt"

func main() {
	//2、变量的初始化，变量声明时同时赋值（一步到位）
	var b int = 10
	b = 20
	b = 80
	b = 90 //赋值，可以使用n次
	fmt.Println("b = ", b)

	//3、自动推导类型，必须初始化，通过初始化的值确定类型(常用）
	c := 30
	//%T打印变量所属的类型
	fmt.Printf("c type is %T\n", c)

	d := 40
	fmt.Println("d=", d)
	//：=，自动推导类型，先声明变量d，再给d赋值为40
	//自动推导，同一个变量名只能使用一次，用于初始化那次
	//前面已经有变量d，不能再新建一个变量d
	d = 50 //只是赋值，是可以的
	fmt.Println("d=", d)

}
