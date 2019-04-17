package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b0
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))

	a = 0

	defer calc("2", a, calc("20", a, b))

	b = 1
}

/*
输出结果为：
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
*/

/*
先明确两个概念：
1、defer是在函数末尾的return前执行，先进吼执行
2、函数调用时int参数发生值拷贝

不管代码顺序如何，defer calc func 中参数b必须先计算
故会在运行到第三行时没执行calc("10", a, b)输出：10 1 2 3
得到值3，将calc("1", 1, 3)存放到延后执行函数队列中。

这里 b = 1 并未被函数使用！！！！*/
