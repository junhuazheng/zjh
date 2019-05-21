package main

import "fmt"

type Myint int

type Element interface{}

var (
	x int
	y Myint
)

func main() {
	x = 100
	y = 100
	list := make([]Element, 2)
	list[0] = x
	list[1] = y
	fmt.Println(list)

	for k, v := range list {
		switch value := v.(type) { //对数据类型进行断言
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", k, value)
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n", k, value)
		case Myint:
			fmt.Printf("list[%d] is an Myint and its value is %v\n", k, value)
		default:
			fmt.Printf("list[%d] is of a different\n", k)
		}
	}
}
