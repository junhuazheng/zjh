/*
Go语言中面对对象是借助struct结构体实现的。
Go语言中虽然没有class关键字来表示类，但却有interface来表示借口。
*/

package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// person := Person{"mike", 18}
// fmt.Println(Person)

/*以上代码用结构体表示了一个人，用面对对象的思维来理解
name和age是两个成员变量，那么我们应该如何编写成员方法呢？
go语言是这样设计的：*/

func (perosn *Person) showInfo() {
	fmt.Printf("My name is %s, age is %d\n", perosn.name, perosn.age)
}

func (perosn *Person) selAge(age int) {
	perosn.age = age
}

func main() {
	perosn := Person{"mike", 18}
	perosn.showInfo()
	perosn.selAge(20)
	fmt.Println(perosn)
}

/*
在关键字func和函数名之间加上一对小括号
写上与之关联的结构体类型的指针类型和变量名
推荐写成指针类型的原因是，函数传递的形参不是指针类型的话，无法覆盖实参。
(个人理解，不使用指针类型的话，只是拷贝）

实际上，此时，showInfo(), serAge()就可以叫做成员方法了。
*/
