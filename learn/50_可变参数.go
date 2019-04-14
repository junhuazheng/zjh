package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, "\n")
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

/*可变参数的函数可以用任何数量的参数来调用。例如，fmt.Println()就是一个常见的可变函数。

这里有一个函数将任意数量的int作为参数。可变参数的函数可以通过单独的参数以通常的方式调用。
如果已经在一个切片中有多个参数，使用 func(slice ...)函数将他们应用到一个可变函数。

Go中的函数的另一个关键方面是他们形成闭包的能力。*/
