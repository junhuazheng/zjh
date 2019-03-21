//匿名函数体可以被赋值：
f := func(data int) {
	fmt.Println("hello", data)
}

//使用f()调用
f(100)

//下面的代码实现对切片的遍历操作，遍历中访问每个元素的操作勇士匿名函数来实现。用户传入不同的
//匿名函数体可以实现对元素不同的遍历操作：

package main

import"fmt"

//遍历切片的每一个元素，通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	
	for _, v := range list {
		f(v)
	}
}

func main() {
	
	//使用匿名珊瑚打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int){
		fmt.Println(v)
	})
}