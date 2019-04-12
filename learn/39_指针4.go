package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
/*使用了2相对应的函数:zeroval()和zeroptr()。
zeroval()函数有一个int参数，因此参数将通过值传递给他。
zeroval将获得ival的靠背，它与调用函数中的值有所不同。

相反，zeroptr有一个*int参数，这就则意味着它需要一个int指针。
函数体重的*iptr代码将指针从存储器地址解引用到该地址处的当前值。
将值分配个取消引用的指针会更改引用地址处的值。

&i语法获取了i变量的存储器地址，即指向i的指针。指针也可以打印。
在main函数中zerobal不会改变i的值，单zeroptr会，以为它有一个对该变量的内存地址的引用。
