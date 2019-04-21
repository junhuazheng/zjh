/*
recover在英文中指的是受伤的痊愈，防止伤口进一步感染。愈合只不过是事后处理。
Golang中的recover知识发生宕机之后的后事处理。

recover用来接收panic触发的宕机
如果panic触发宕机，传给panic的任意类型的参数，recover会接收到这个参数
recover获取到值之后才知道发生了宕机；
相反，如果程序中的recover没有获取到值，则代表没有发生宕机，那么recover的值就为nil
通过这个可以来进一步处理后事。

前面已经提到了panic的时候，已经说了，一旦发生宕机，其后的代码是不执行的
但是会调用位于panic代码所在的哪一行之前的defer延迟函数
所以这个特性就决定recover应该用在defer函数中
否则一旦发生宕机，除了defer延迟函数中的语句还能执行外，其他的语句都是不能执行的。

如果触发宕机，panic的错误信息会显示，如果有recover时：
则信息会被recover截获，于是错误信息就不会显示，转而进行下一步的操作
*/

package main

import "fmt"

func main() {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("触发了宕机", info)
		} else {
			fmt.Println("程序正常退出")
		}
	}()
	fmt.Println("aaaaaa")
	fmt.Println("bbbbbb")
	panic("fatal error")
	fmt.Println("cccccc")

	defer func() {
		fmt.Println("ddddddd")
	}()

}
