//指针类型的接收器由一个结构体的指针组成，更接近于面向对象中的this或者self
//由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的
//使用结构体定义一个属性（Property），为属性添加SetValue()方法以封装设置属性的过程。
//通过属性的Value()方法可以重新获得属性的数值，通过SetValue()方法的调用，可以达成修改属性的效果。
package main

import "fmt"

//定义属性结构，拥有一个整型的成员变量
type Property struct {
	value int //属性值
}

//设置属性值（定义属性值得方法）
func (p *Property) SetValue(v int) {

	//设置属性值方法的接收器为指针。因此可以修改成员值，修改p的成员变量，即便退出方法也有效。
	p.value = v
}

//取属性值（定义获取值得方法）
func (p *Property) Value() int {
	return p.value
}

func main() {

	p := new(Property)

	p.SetValue(100) //设置值，此时成员变量为100.

	fmt.Println(p.Value()) //获取成员变量
}
