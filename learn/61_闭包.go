/*
函数列表使用不当
*/
package main

import "fmt"

func test() []func() {
	var s []func()

	for i := 0; i < 3; i++ {
		s = append(s, func() { //将多个函数添加到列表
			fmt.Println(&i, i)
		})
	}

	return s //返回匿名函数列表
}

func main() {
	for _, f := range test() { //执行所有匿名函数
		f()
	}
}

/*输出结果为：
0xc000010090 3
0xc000010090 3
0xc000010090 3

每次append操作仅将匿名函数放入列表中，但并未执行，并引用的变量都是i
随着i的改变匿名函数中的i也在改变，所以当执行这些函数时
他们读取到的都是环境变量i最后一次的值。

解决方法就是每次复制变量i然后传到匿名函数中，让闭包的环境变量不同
*/

//修改方法1：
func test() []func() {
	var s []func()

	for i := 0; i < 3; i++ {
		x := i
		s = append(s, func() { //将多个函数添加到列表
			fmt.Println(&x, x)
		})
	}

	return s //返回匿名函数列表
}

func main() {
	for _, f := range test() { //执行所有匿名函数
		f()
	}
}

/*输出结果为：
0xc00004a058 0
0xc00004a070 1
0xc00004a078 2*/

//修改方法2
var x int = 1

func main() {
	y := func() int {
		x += 1
		return x
	}()
	fmt.Println("main:", x, y)
}

//输出结果为： main: 2 2
