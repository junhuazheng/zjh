//初始化定义的类型
//先声明再赋值

package main

import "fmt"

type Chairman struct {
	Name string
	age  int
}

func main() {
	var c Chairman
	c.Name = "习大大"
	c.age = 64
	fmt.Println(c)
}
