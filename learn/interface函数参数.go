//模拟实现”fmt.Println()”的Stringer方法
/*
type Stringer interface {
	String() string
}
*/

package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name   string
	age    int
	string //定义住址，这是匿名字段
}

func (s Student) String() string {
	return "My name is " + s.Name + ", I am " + strconv.Itoa(s.age) + " years old . I live in " + s.string
}

func main() {
	zjh := Student{"郑俊华", 99, "广州"}
	fmt.Println("This people is :", zjh)
}
