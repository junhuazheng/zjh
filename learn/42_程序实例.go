package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string

	fmt.Scanln(&input)
	fmt.Println("done")
}

//两个函数调用现在在不同的goroutine中异步进行，所以执行到这里。Scanln代码要求我们在退出前按一个键
