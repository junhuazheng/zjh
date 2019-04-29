package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	
	fmt.Printf("%+v", p) //如果值是一个结构体，%+v的格式化输出内容将包括结构体的字段名
	
	fmt.Printf("%#v", p) //输出这个值得Go语法表示
	
	fmt.Printf("%T\n", p)
	
	fmt.Printf("%t\n", true)
	
	fmt.Printf("%d\n", 123)
	
	fmt.Printf("%b\n", 123) //输出二进制表达式
	
	fmt.Printf("%c\n", 123) //输出给定整数的对应字符
	
	fmt.Printf("%x\n", 123) //%x提供十六进制编码
	
	fmt.Printf("%f\n", 123.1) //浮点型
	
	fmt.Printf("%s\n", "\"string"\") //字符串
	
	//Printf她通过os.Stdout输出格式化的字符串
	//Sprintf则格式化并返回一个字符串而不带任何输出。
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s))
	
	//使用Fprintf来格式化并输出到io.Writers而不是os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}