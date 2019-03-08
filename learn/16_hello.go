package main

import "fmt"

func main() {
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t = ", t)

	//自动推到类型
	t2 := 3.3 + 4.4i
	fmt.Printf("t2 type is %T\n", t2)

	//通过内建函数，取实部和虚部
	fmt.Println("real(t2) = ", real(t2), ", imag(t2) = ", imag(t2))
	fmt.Println("real(t) = ", real(t))

	a := 10
	b := "abc"
	c := 'a'
	d := 3.14
	//%T操作变量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

	//%d 整形格式
	//%s 字符串格式
	//%c 字符格式
	//%f 浮点型格式
	fmt.Printf("a = %d, b = % s, c = %c, d = %f\n", a, b, c, d)
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d) //%v自动匹配格式化输出
}
