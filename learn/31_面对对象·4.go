/*
什么是对象？
对象，指的是客体。所谓客体是指客观存在的对象实体和主观抽象的概念。
简单理解就是，抽象一个同游多重属性的客体，将共有属性抽离出来为一个类，以便实现定义多个客体的功能。

面对对象有哪些特征？
面对对象通常包括三个特征 继承，封装和多态：
1、继承：由子类继承父类的属性/数据/方法等。
2、封装：以最简单的函数形式将方法展示出去，而不需要使用者知道方法内有什么，
由什么实现，类似黑盒子，只需知道怎么用，无需知道为什么。

在Go语言中，通常使用接口的方式完成封装；
3、多态：一中方法的多种变现形式，可以看作是封装后的方法的结合
根据使用场景，自动分发到某具体方法中，即一个同样的函数对于不同的对象可以具有不同的实现。

为什么要使用面对对象？
1、面对对象是为了了解系统的可维护性，可扩展性，可重用性。
2、简单理解：以对象方法代替过程完成时限，方便以后修改及复用。

以简单计算器为例，展示Go语言如何实现面对对象。
*/

package main

import (
	"fmt"
)

type BaseNum struct {
	num1 int
	num2 int
}

type Add struct {
	BaseNum
}

type Sub struct {
	BaseNum
}

func (a *Add) Opt() (value int) {
	return a.num1 + a.num2
}

func (s *Sub) Opt() (value int) {
	return s.num1 - s.num2
}

type Opter interface {
	OPt() int
}

func MultiState(o Opter) (value int) {
	value = o.OPt()
	return
}

func main() {
	var a Add = Add{BaseNum{2, 3}}

	value := a.Opt()

	var i Opter
	i = &a
	value = i.OPt()

	i = &a
	value = MultiState(i)
	fmt.Println(value)
}
