//结构体内嵌模拟类的继承
//在面向对象思想中，实现对象关系需要使用“继承”特性。例如人类不能飞行，鸟类可以飞行。人类和鸟类都可以继承自可行走类，但只有鸟类继承自飞行类

//面对对象的设计原则中建议对象最好不要使用多重继承，有些面向对象语言从语言层面就禁止了多重继承

//Go语言的结构体内嵌特性就是一种组合特性，使用组合特性可以快速构建对象的不同特性。

//下面代码使用Go语言的结构体内嵌视线对象特性组合
//人和鸟的特性：

package main

import "fmt"

type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

type Walkable struct{}

func (f *Walkable) Walk() {
	fmt.Println("can walk")
}

type Human struct {
	Walkable
}

type Bird struct {
	Walkable
	Flying
}

func main() {

	b := new(Bird)
	fmt.Println("Bird: ")
	b.Fly()
	b.Walk()

	h := new(Human)
	fmt.Println("Human: ")
	h.Walk()
}

//使用Go语言的内嵌结构体实现对象特性，可以自由地在对象中增、删、改各种特性。Go语言会在编译时检查能否使用这些特性
