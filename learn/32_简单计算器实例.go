/*
1、抽象类型：计算器，可以抽象为两个数字，一个运算符和一个结果返回值；
父类：两个数字
子类：继承父类
子类的方法：做出计算并输出结果返回值
2、定义方法，对不同的运算符返回不同的运算结果
3、封装：定义借口，将子类方法进行封装
4、多态：定义多态，并将封装好的接口作为形参，实现多态；
！！！可以简单理解为，以接口作为形参的函数。
*/

package main

import (
	"fmt"
)

//分析实现过程，进行抽象化：两个数字，一个运算符，一个结果返回值。
type BaseNum struct {
	num1 int
	num2 int
} //BaseNum 即为父类型名称

type Add struct {
	BaseNum
} //加法子类，定义加法子类的主要目的，是为了定义对应子类的方法，

type Sub struct {
	BaseNum
} //减法子类

//定义子类方法，实现运算及返回值。
func (a *Add) Opt() (value int) {
	return a.num1 + a.num2
} //加法的方法实现

func (s *Sub) Opt() (value int) {
	return s.num1 - s.num2
} //减法的方法实现

//注意，这里的方法名称是一样的，这样才能使用接口进行归纳。
type Opter interface { //接口定义
	Opt() int //封装，归纳子类方法，注意此处需要加上返回值，不然没办法输出返回值（因为方法中使用了返回值）
}

//定义多态
func MultiState(o Opter) (value int) { //多态定义，可以简单理解为接口作为形参的函数。
	value = o.Opt()
	return
}

//主函数及调用
func main() {
	var a Add = Add{BaseNum{2, 3}}

	//使用Add对象方法
	value := a.Opt()

	//使用接口
	var i Opter
	i = &a
	value = i.Opt()

	//使用多态
	i = &a
	value = MultiState(i)

	//输出
	fmt.Println(value)
}
