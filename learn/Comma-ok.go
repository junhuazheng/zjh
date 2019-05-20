/*interface的变量里面可以存储任意类型的树枝（该类型实现了interface）。
那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？
有两种方法：“Comma-ok断言”和“switch测试“

1、Comma-ok断言*/

package main

import (
	"fmt"
	"strconv"
)

type Element interface{}

const (
	pai = 3.14
)

type Student struct {
	Name string
	age  int
}

func (p Student) String() string {
	return "(name:" + p.Name + "-age:" + strconv.Itoa(p.age) + "years old)"
}

func main() {
	list := make([]Element, 4)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Student{"Zhengjunhua", 99}
	list[3] = pai

	for index, element := range list { //判断里面的每一个元素属于哪一种类型
		if value, ok := element.(int); ok { //判断当前的数据类型是否为“int”类型
			fmt.Printf("list[%d] is an int an its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is an string an its value is %s\n", index, value)
		} else if value, ok := element.(Student); ok {
			fmt.Printf("list[%d] is an Student an its value is %v\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type!", index)
		}
	}
}
