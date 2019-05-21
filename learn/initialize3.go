//按照key和value的方式初始化

package main

import "fmt"

type Chairman struct {
	Name string
	age  int
}

func main() {
	var c Chairman
	c = Chairman{
		age:  64,
		Name: "习大大",
	}

	fmt.Println(c)
}
