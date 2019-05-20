/*
type Interface interface {
	sort.Interface
	Push(x interface{}) //add x as element Len()
	Pop() interface{}   //remove and return element Len() - 1
}

看到官方使用的嵌入interface方法可以联想到之前的你名字段。
就像struct的匿名字段，那么相同的逻辑引入到interface里面是什么效果呢？
如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式
的包含了interface1里面的method。

下面这个例子自定义实现了嵌入interface
*/

package main

import "fmt"

type Employee struct {
	Name   string
	age    int
	salary int
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) GetAge() int {
	return e.age
}

func (e *Employee) GetSalary() int {
	return e.salary
}

func (e *Employee) Help() {
	fmt.Println("Don't ask me, I won't tell you")
}

type BasicInformation interface {
	GetName() string
	GetAge() int
}

type SpecificInformation interface {
	BasicInformation
	GetSalary() int
	Help()
}

func main() {
	zjh := Employee{
		Name:   "郑俊华",
		age:    99,
		salary: 80000,
	}

	fmt.Println("zjh is: ", zjh)
	zjh.Help()
	fmt.Println("zjh.name = ", zjh.GetName())
	fmt.Println("zjh.age = ", zjh.GetAge())
	fmt.Println("zjh.salary = ", zjh.GetSalary())

	var zhengjunhua SpecificInformation = &zjh

	switch zhengjunhua.(type) {
	case nil:
		fmt.Println("空接口")
	case BasicInformation:
		fmt.Println("BasicInformation 接口")
	default:
		fmt.Println("位置接口")
	}
}
