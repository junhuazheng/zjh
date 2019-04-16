package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) //err就是panic传入的参数
		}
		fmt.Println("打印前")
	}()

	defer func() { //必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) //err就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) //err就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

/*
输出结果为：
触发异常
打印后
打印中
打印前
*/

/*总结：
1、defer函数属延迟执行，延迟到调用者函数执行return命令前被执行。
多个defer之间按LIFO先进后出顺序执行。

2、Go中可以跑出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

3、如果同时有多个defer，那么异常会被最近的recover()捕获并正常处理。
*/
