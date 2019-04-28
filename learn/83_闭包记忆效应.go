/*
被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量
变量会随着闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应
*/

package main

import (
	"fmt"
)

func Acc(value int) func() int {

	//返回一个闭包
	return func() int {

		value++

		return value
	}
}

func main() {

	//创建一个累加器，初始值为1
	acc := Acc(1)

	//累加1并打印
	fmt.Println(acc())

	fmt.Println(acc())

	//打印累加器的函数地址
	fmt.Printf("%p\n", acc)

	acc2 := Acc(10)

	fmt.Println(acc2())

	fmt.Printf("%p\n", acc2)
}
