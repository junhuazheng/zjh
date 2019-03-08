package main

import "fmt"

func main() {
	var str1 string //声明变量
	str1 = "abc"
	fmt.Println("str1 = ", str1)

	//自动推到类型
	str2 := "mike"
	fmt.Printf("str2 type is %T\n", str2)

	//内建函数，len（）可以测字符串长度，有多少个字符
	fmt.Println("len(str2) = ", len(str2))

	var ch byte
	var str string

	//字符
	//1、单引号
	//2/字符，往往都只有一个字符，转义字符除外'\n'
	ch = 'a'
	fmt.Println("ch = ", ch)
	//字符串
	//1、双引号
	//2、字符串由一个或多个字符组成
	//3、字符串都是隐藏了一个结束符，'0'
	str = "a" // 由'a'和''\0'组成了一个字符串
	fmt.Println("str = ", str)

	str = "hello go"
	//只想操作字符串的某个字符，从0开始操作
	fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])
}
