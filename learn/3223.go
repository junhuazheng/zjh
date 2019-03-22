package main

import "fmt"

//使用type MyInt int 将int定义为自定义的MyInt类型
type MyInt int

//为MyInt类型添加IsZero()方法。该方法使用了(m MyInt)的非指针接收器。数值类型没有必要使用指针接收器
func (m MyInt) IsZero() bool {
	return m == 0
}

func (m MyInt) Add(other int) int {
	return other + int(m) //由于m的类型是MyInt类型，但其本身是int类型，因此可以将m从MyInt类型转换为int类型再进行计算
}

func main() {

	var b MyInt

	//调用b的IsZero()方法。由于使用非指针接收器，b的值会被复制进入IsZero()方法进行判断
	fmt.Println(b.IsZero())

	b = 1

	fmt.Println(b.Add(2))
}
