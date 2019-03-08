package main

import "fmt"

func main() {
	var a int //声明变量
	fmt.Printf("请输入变量a: ")

	//阻塞等待用户的输入
	//fmt.Scanf("%d", &a)
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
