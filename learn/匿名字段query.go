//匿名字段的查询操作
package main

import "fmt"

type Student struct {
	Name       string
	age        int
	speciality string
	int        //内置类型作为匿名字段
}

type Classroom struct {
	Student //默认Classroom就包含了Student的所有字段。
	area    int
	name    string
}

func MyEcho(p Classroom) {
	fmt.Printf("学生的姓名是: [%s]\n", p.Student.Name)
	fmt.Printf("学生的身高是: [%d]\n", p.Student.int)
	fmt.Printf("学生的爱好是: [%s]\n", p.Student.speciality)
	fmt.Printf("学生的年龄是: [%d]\n", p.Student.age)
	fmt.Printf("教室的面积是: [%d]\n", p.area)
	fmt.Printf("教室的名称是: [%s]\n", p.name)
}

func main() {
	zhengjunhua := Classroom{
		Student{
			Name:       "郑俊华",
			age:        18,
			speciality: "Baseball",
			int:        175,
		},
		100,
		"Golang学习",
	}
	MyEcho(zhengjunhua)
}
