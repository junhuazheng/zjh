package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() { //必须要先声明defer，否则recover()不能捕捉到panic异常

		if err := recover(); err != nil {
			fmt.Println(err) //err就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

/*
输出结果为：
打印后
触发异常
打印中
打印前
*/
