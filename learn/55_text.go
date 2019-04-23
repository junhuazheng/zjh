package main

import "fmt"

func Text(a bool) (err error) {
	if a {
		fmt.Println("true111")
	}
	fmt.Println(err)
	return err
}

func main() {
	var a bool = false
	var b bool = true
	Text(a)
	Text(b)

}

/*
输出结果为：true111
往函数Text输入false，将直接返回
往函数Text输入true，打印“true111”然后返回
*/
