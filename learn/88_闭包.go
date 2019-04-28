package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

/*输出结果为：
1
2
3
1*/

/*
通过闭包来使用匿名函数。匿名函数在你想定义一个不需要命名的内联函数是很实用的。

这个intSeq函数返回另一个在intSeq函数体内定义的匿名函数。这个返回的函数使用闭包的方法 隐藏变量i。

我们调用intSeq函数，将返回值（也是一个函数）赋给nextInt。
这个函数的值包含了自己的值i，这样在每次调用nextInt时都会更新i的值。
*/
