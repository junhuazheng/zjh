/*
面对对象编程，分为几个步骤？

面对对象是一种思想，他让我们在分析和解决问题时，把思维和重点转向现实中的客体中来，
然后通过UML工具清理这些可以之间的联系，最后用面向对象的语言实现这种客体以及客体之间的联系。
它分为面向对象的分析（OOA），面向对象的设计（OOD），面向对象的编程实现（OOP）三个大的步骤。

1、首先是分析需求，先不要思考怎么用程序来实现它，先分析需求中稳定不变的客体都是些什么。这些客体之间的关系是什么。

2、把第一步分析出来的需求，通过进一步扩充模型，变成可实现的、符合成本的、模块化的、低耦合高内聚的模型。

3、使用面向对象的视线模型。
*/
/*
在上面的实例中，我们提到了运算符，并将运算符和输入值和输出值并列在一块，这是为什么？
因为我们可以通过实现模型来完成更加简洁的写法：
下面实例使用工厂模式来解决计算器的问题：
*/
package main

import "fmt"

/*
实例：面向对象的计算器实现
1、定义父类
2、定义子类，以及子类的方法 运算实现
3、定义接口，归纳子类方法
4、定义空类，定义空类的方法，即 工厂模式，将 运算符 与 数值 分开出来，以运算符来分发方法，方便调用
5、定义一个多态，将接口管，方便调用
6、主函数，初始化，调用工厂模式，进行验证。
*/

//父类
type BaseNum struct {
	num1 int
	num2 int
}

//加法子类
type Add struct {
	BaseNum
}

//减法子类
type Sub struct {
	BaseNum
}

//子类方法
func (a *Add) Opt() int {
	return a.num1 + a.num2
}

func (s *Sub) Opt() int {
	return s.num1 - s.num2
}

//定义接口，即封装
type Opter interface {
	Opt() int
}

//定义多态
func MultiState(o Opter) int {
	value := o.Opt()
	return value
}

func (f *Factory) FacMethod(a, b int, operator string) (value int) {
	var i Opter
	switch operator {
	case "+":
		var AddNum Add = Add{BaseNum{a, b}}
		i = &AddNum
	case "-":
		var SubNum Sub = Sub{BaseNum{a, b}}
		i = &SubNum
	}

	//接口实现 : value = i.Opt()
	value = MultiState(i) //多态实现
	return
}

func main() {
	var a Factory
	value := a.FacMethod(20, 3, "-")
	fmt.Println(value)
}
