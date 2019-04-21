/*函数特性
init函数作为Go的两个保留函数之一（另一个是main函数）有很多特殊性;
1、init函数没有入参和返回值
2、init函数不能被显式调用
3、一个package中可以包含多个init函数
4、一个go文件中可以包含多个init函数
*/
package main

import "fmt"

var WhaiIsThe = AnSwerToLife()

func AnSwerToLife() int {
	return 43
}

func init() {
	WhaiIsThe = 0
}

func main() {
	if WhaiIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}

/*
程序的初始化都从pkg main开始，如果main引入了其他package，
在编译的时候就会将他们一次导入。

初始化是在一个goroutine中进行的，init函数可以引入新的goroutine
但只有在当前init函数返回后才会执行下一个init。
*/
