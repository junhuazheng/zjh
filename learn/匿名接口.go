/*   我们可以在一个结构体中定义一个匿名字段，这个匿名字段可以使内置的也可以是我
们自己定义的，而interface是一种特殊的数据类型，因为golang认为所有的数据类型
都实现了空接口，也就是说所有数据都是空interface的子集。换句话说，我们可以说一
个空interface是可以接受任何类型的数据的。通过这点，他会给我们很多启发，我们可以
通过interface来接受任何类型的参数，也可以通过interface来返回任何类型的参数。
下面来看一个匿名interface的使用案例：*/
package main

import "fmt"

type Employee struct {
	Name   string
	age    int
	salary int
	gender string
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

func (e *Employee) GetGender() string {
	return e.gender
}

type MyEmployee struct {
	GenInfo interface { //匿名接口可以被用作变量或者结构属性类型，我们定义了一个“GenInfo”的匿名接口，里面可以额存储各种数据属性。
		GetGender() string
		GetSalary() int
		GetAge() int
		GetName() string
	}
}

func main() {
	zjh := MyEmployee{&Employee{
		Name:   "郑俊华",
		age:    99,
		salary: 80000,
		gender: "boy",
	}}
	fmt.Println("姓名:", zjh.GenInfo.GetName())
	fmt.Println("年龄:", zjh.GenInfo.GetAge())
	fmt.Println("性别:", zjh.GenInfo.GetGender())
	fmt.Println("期望薪资:", zjh.GenInfo.GetSalary())
}
