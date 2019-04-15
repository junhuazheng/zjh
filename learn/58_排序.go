/*Go语言的sort包实现了内置和用户定义类型的排序。

我们先看看内置的排序。排序方法特定于内置类型；
这里是一个字符串的例子，请注意，排序是就地排序，因此他会更改给定的切片，并且不返回新的切片。

也可以使用sort来检查切片是否已经按照排序顺序。

运行程序打印排序的字符串，以及int数据值和true作为AreSorted测试的结果。*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)

	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)

	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints) //检查切片是否已经按照排序顺序
	fmt.Println("Sorted: ", s)
}
