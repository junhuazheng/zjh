package main

import "fmt"

func main() {
	//声明变量
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)

	//自动推到类型
	f2 := 3.14
	fmt.Printf("f2 type is %T\n", f2) //f2 type is float64

	//float64存储小数比float32更准确一点

	//字符类型:ascii码

}
