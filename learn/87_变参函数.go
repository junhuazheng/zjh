/*
可变参数函数，可以用任意数量的参数调用。例如，fmt.Println是一个常见的变参函数。
*/
package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

/*
如果你的slice已经有了多个值，想把他们作为变参使用，你要这样调用func(slice...)
*/
