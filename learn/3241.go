//Go语言内嵌结构体成员名字冲突

//嵌入结构体内部可能拥有相同的成员名，成员名重名时会发生什么？
package main

import "fmt"

type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

func main() {
	c := &C{}      //实例化C结构体
	c.A.a = 1      //按常规方法，访问嵌入结构体A中的a字段，并赋值1
	fmt.Println(c) //可以正常输出实例化C结构体
}

//接下来，将c.A.a = 1修改为下面代码：
// func main() {
// 	c := &C{}
// 	c.a = 1
// 	fmt.Println(c)
// }

//此时再编译运行，编译器报错：
//.\main.go.22:3:ambiguous selector c.a
//编译器告知C的选择器a引起歧义，也就是说你，编译器无法界定将1赋给C中的A还是B里的字段a

//使用内嵌结构体时，Go语言的编译器会非常智能地体型我们可能发生的歧义和错误
