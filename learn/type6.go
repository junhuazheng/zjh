// 匿名结构体定义格式和初始化方法：
// ins := struct {
// 	字段1 字段类型1
// 	字段2 字段类型2
// 	...
// }{
// 	初始化字段1:字段1的值,
// 	初始化字段2:字段2的值,
// 	...
// }

package main

import "fmt"

func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
}

func main() {

	msg := &struct {
		id   int
		data string
	}{
		1024,
		"world",
	}

	printMsgType(msg)
}
