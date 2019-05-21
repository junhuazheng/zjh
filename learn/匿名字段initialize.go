/*结构体的字段可以只有类型名，而没有字段名，这种字段称为匿名字段。
匿名字段可以是一个结构体、切片等符合类型，也可以使int这样的简单类型。
但建议不要把简单类型作为你名字段。golang中每种类型只能有一个匿名域。
可以用来实现oop中的继承。

匿名字段的初始化操作*/

package main

import "fmt"

type Student struct {
	name       string
	age        int
	speciality string
	int        //内置类型作为匿名字段
}

type Classroom struct {
	Student //默认Classroom就包含了Student的所有字段。
	area    int
	name    string
}

func main() {
	zhengjunhua := Classroom{
		Student{
			name:       "郑俊华",
			age:        18,
			speciality: "Baseball",
			int:        175,
		},
		100,
		"Golang学习之路",
	}
	fmt.Println(zhengjunhua)
}
