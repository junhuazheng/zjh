//按照位置初始化

package main

import "fmt"

type Chairman struct {
	Name string
	age  int
}

func main() {
	var c Chairman
	c = Chairman{
		"习大大",
		64,
	}

	fmt.Println(c)
}
