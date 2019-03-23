//使用Go语言结构体内嵌写法：
package main

import "fmt"

type BasicColor struct {
	R, G, B float32
}

type Color struct {
	BasicColor //将BasicColor结构体嵌入到Color结构体中，BasicColor没有字段名只有类型，这种写法叫做结构体内嵌
	Alpha      float32
}

func main() {

	var c Color
	//可以直接对Color的R、G、B成员进行设置，编译器通过Color的定义知道R、G、B成员雷子BasicColor内嵌的结构体
	c.R = 1
	c.B = 0
	c.G = 1

	c.Alpha = 1

	fmt.Println("%+v", c)
}
