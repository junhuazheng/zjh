package main

import "fmt"

type class struct {
}

//为结构体添加一个Do()方法，参数为整型。这个方法的功能呢过是打印提示和输入的参数值
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

//声明一个普通函数，参数也是整型，功能是打印提示和输入的参数值
func funcDo(v int) {
	fmt.Println("call function do:", v)
}

func main() {
	var delegate func(int) //声明一个delegate()函数，类型为func(int)，与funcDo和class的Do()方法的参数一致

	c := new(class)

	delegate = c.Do //将c.Do作为值付给delegate变量

	delegate(100) //调用delegate()函数，传入100参数。此时会调用c实例的Do方法

	delegate = funcDo //将funcDo赋值给delegate

	delegate(100) //调用delegate()，传入100的参数。此时会调用funcDo()方法
}

//这段代码能运行的基础在于：
//无论是普通函数还是结构体的方法，只要他们的签名一致，与他们签名一致的函数变量就可以保存普通函数或是结构体方法
