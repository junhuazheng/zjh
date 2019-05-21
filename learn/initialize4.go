//自定义struct案例及其用法示例

package main

import "fmt"

type Student struct {
	name string
	age  int
}

//将定义的struct类型传递进来比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
func AgeConstrast(p1, p2 Student) (Student, int) {
	x := p1.age - p2.age
	y := p2.age - p2.age
	if p1.age > p2.age {
		return p1, x
	} else {
		return p2, y
	}
}

func main() {

	var zhengjunhua Student

	zhengjunhua.name, zhengjunhua.age = "郑俊华", 18

	cat := Student{age: 25, name: "猫"} //按照key：value的方式初始化赋值

	Leader := Student{"马云", 55}

	zjh_cat, a := AgeConstrast(cat, zhengjunhua)
	zjh_my, b := AgeConstrast(Leader, zhengjunhua)
	cat_my, c := AgeConstrast(Leader, cat)

	fmt.Printf("当[%s]和[%s]在一起的时候，[%v]比他大[%d]岁\n", zhengjunhua.name, cat.name, zjh_cat.name, a)
	fmt.Printf("当[%s]和[%s]在一起的时候，[%v]比他大[%d]岁\n", zhengjunhua.name, Leader.name, zjh_my.name, b)
	fmt.Printf("当[%s]和[%s]在一起的时候，[%v]比他大[%d]岁\n", cat.name, Leader.name, cat_my.name, c)
}
