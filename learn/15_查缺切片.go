package main

import "fmt"

func main() {
	var hRise [30]int

	for i = 0; i < 30; i++ {
		hRise[i] = i + 1
	}

	var strL []string

	var strK []int

	var strJ = []int{}
	
	a := make([]int, 2, 5) //这个5 是预分配的元素数量，这值只是能提前分配空间，降低多次分配空间造成的性能问题。
	
	fmt.Println(a)
	fmt.Println(len(a))

	fmt.Println(strL, strK, strJ)
	fmt.Println(strJ = nil) //结果为false，strJ已经被分配到了内存，因此不为空

	fmt.Println(hRise[10:15])
	fmt.Println(hRise[20:])
	fmt.Println(hr[:2])
}

//重置切片，清空拥有的元素
// a := []int{1, 2, 3}
// fmt.Println(a[0:0])

//添加切片
// team := []string{"pi", "fi", "ki"}
// avr car []string
// car = append(car, team...)
