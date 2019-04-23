package main

import (
	"fmt"
	"sync"
)

func p(i int) {
	fmt.Println(i)
}

func main() {
	a := []int{1, 2, 3}

	for _, i := range a {
		fmt.Println(i)
		defer p(i)
	}
	main1()
}

/*
输出结果为：
1
2
3
3
2
1
这里就是普通的函数调用，defer由于是后进先出，所以执行变成3,2,1
*/
