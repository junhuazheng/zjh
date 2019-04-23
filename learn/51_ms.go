package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		//loop:
		println(i)
	}
	//goto loop
	main1()
}

/*goto
goto不能跳转到其他函数或者内层代码
输出结果为:
goto loop jumps into block starting at
*/

func main1() {
	type MyInt1 int
	type MyInt2 = int

	var i int = 9
	var i1 MyInt1 = i //error
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

//输出结果为:cannot use i(type int) as type MyInt1 in assignment

/*
我们给予一个类型创建一个新类型，称之为defintion
给予一个类型创建一个别名，称之为alias，这就是他们最大的区别

type MyInt1 int
type MyInt2 = int
第一行代码是基于基本类型int创建了新类型MyInt1,
第二行是创建的一个int的类型别名的MyInt2
注意，类型别名的定义是 =

var i int = 0
var i1 MyInt1 = i
var i2 MyInt2 = i
fmt.Println(i1, i2)
第二行吧一个int类型的变量i，赋值给MyInt1类型的变量i1会被提示编译错误:类型无法转换。
但是第三行吧int类型的变量i，赋值给MyInt2类型的变量i2就可以，不会提示错误。

从这里可以看出，这两种定义方式的不同，因为Go是强类型语言
所以类型之间的转换必须强制转换，因为int和MyInt1是不同类型，所以这里会报编译错误。

但是因为MyInt2只是int的一个别名，本质上还是一个int类型，所以可以直接赋值，不会有问题。
*/
