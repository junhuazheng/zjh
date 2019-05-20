package main

import (
	"fmt"
	"strconv"
)

type Element interface{}

type Student struct {
	Name string
	age  int
}

func (s Student) String() string {
	return "(name:" + s.Name + "-age:" + strconv.Itoa(s.age) + "years old)"
}

func main() {
	list := make([]Element, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Student{"Zhengjunhua", 99}

	for k, v := range list {
		switch value := v.(type) { //element.(type)语法不能再switch外的任何逻辑里面使用，如果要在switch外面判断一个类型就用 comma-ok。
		case int:
			fmt.Printf("list[%d] is an int an its value is %d\n", k, value)
		case string:
			fmt.Printf("list[%d] is an string an its value is %s\n", k, value)
		case Student:
			fmt.Printf("list[%d] is an Student an its value is %v\n", k, value)
		default:
			fmt.Printf("list[%d] is of a different\n", k)
		}
	}

}
