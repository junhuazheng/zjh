package main

import "fmt"

func main() {
	//给int64起一个别名叫bigint
	type bigint int64

	var a bigint //等价于var 啊int64
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d, ch = %c\n", b, ch)

	fmt.Println("4 > 3 结果：", 4 > 3)
	fmt.Println("4 != 3 结果: ", 4 != 3)

	fmt.Println("!(4 > 3) 结果：", !(4 > 3))
	fmt.Println("!(4 != 3) 结果：", !(4 != 3))

	//&& 与， 并且， 左边右边都为真，结果才为真， 其他都为假
	fmt.Println("true && true 结果：", true && true)
	fmt.Println("true && false 结果：", true && false)

	//||, 或者， 左边右边都为假， 结果才为假， 其他都为真
	fmt.Println("true || false 结果：", true || false)
	fmt.Println("false || false 结果：", false || false)

}
