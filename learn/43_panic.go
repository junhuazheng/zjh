/*panic
关键字panic的作用是制造一次宕机，宕机就代表程序运行终止，但是已经“生效”的延迟函数仍会执行
（即已经压入栈的defer延迟函数，panic之前的）

Go程序设计语言中这样提到：
如果碰到“不可能发生的”的状况，宕机是最好的处理方式。
*/

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("aaaaa")
	}()

	fmt.Println("bbbbb")
	fmt.Println("ccccc")
	panic("hahahaha")
	fmt.Println("ddddd")
	defer func() {
		fmt.Println("eeeeeee")
	}()
}

/*
执行程序，会将第一个defer延迟函数“入栈”，然后输出“bbbbb”，“ccccc”
此时使用panic来出发一次宕机
panic接受一个任意类型的参数会将该字符串输出，用作提示信息，之后的代码不再执行
第二个defer延迟函数也不会“入栈“，不会执行
*/
