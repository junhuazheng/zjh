package main

import "fmt"

func main() {
	var str1 string
	str1 = "abc"
	fmt.Println("str1 = ", str1)
	fmt.Printf("%v\n", str1)
	fmt.Printf("str1 type is %T\n", str1)
	fmt.Println("len(str1) = ", len(str1))

	var ch byte
	fmt.Println("ch = ", ch)
	ch = 'a'
	fmt.Println("ch = ", ch)

	var str string
	str = "a"
	fmt.Println("str = ", str)

	var a int
	fmt.Printf("请输入变量：")
	fmt.Scan(&a)
	fmt.Println("a = ", a)

	var t complex128
	t = 2.1 + 3.5i
	fmt.Println("t = ", t)

	t2 := 2.5 + 5.6i
	fmt.Printf("t2 type is %T\n", t2)
	fmt.Println("real(t2)", real(t2), ",imag(t2)", imag(t2))
}
