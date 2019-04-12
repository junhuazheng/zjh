package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(person{name: "Ann", age: 1})

	s := person{name: "San", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51 //指针自动取消引用，相当于s.age = 50
	fmt.Println(sp.age)

}
/*Go语言中的结构是字段的类型集合。它们可用于将数据分组到表单记录。

person结构类型拥有name和age两个字段。创建一个新的结构体。
并且可以在初始化结构时命名字段。忽略的字段将为零值。
&符号为前缀将产生一个指向struct的指针。

使用点(.)来访问结构体中的字段。还可以使用点来访问指针字段-指针将被自动取消引用。结构体是可变的。

