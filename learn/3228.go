//Go语言可以将类型和方法与普通函数视为一个概念，从而简单化方法和函数混合作为回调类型时的复杂性
//调用者无需关心谁来支持调用，系统会自动处理是否调用普通函数或类型的方法。

//先了解Go语言是如何将方法与函数视为一个概念，接着会实现一个事件系统，是事件系统能有效地将事件触发与相应两端代码解耦

//方法和函数的统一调用
//让一个结构体的方法(class.Do)的参数和一个普通函数（funcDo）的参数完全一致，也就是方法与函数的签名一致。
//然后使用与他们签名一致的函数变量（delegate)分别赋值方法与函数，接着调用他们，观察效果
package main

import "fmt"

type class struct {
}

func (c *class) Do(v int) {

	fmt.Println("call metod do :", v)
}

func funcDo(v int) {
	fmt.Println("call function do:", v)
}

func main() {

	var delegate func(int)

	c := new(class)

	delegate = c.Do

	delegate(100)

	delegate = funcDo

	delegate(100)
}
