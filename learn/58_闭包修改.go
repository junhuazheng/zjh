package main

import "fmt"

/*闭包改写程序：*/
func main() {
	a := []int{1, 2, 3}
	for _, i := range a {
		fmt.Println(i)
		defer func() {
			fmt.Println(i)

		}()
	}
}

/*
输出结果为：
1
2
3
3
3
3
闭包里的非传递参数外部变量值是传引用的，在闭包函数里那个i就是外部非闭包函数自己的参数。
所以是相当于引用了外部的变量，i的值执行到第三次是3，闭包是地址引用所以打印了3次i地地址指向的值
所以是3， 3， 3
*/
