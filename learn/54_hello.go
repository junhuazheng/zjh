package main

import "fmt"

func MyFunc() {
	a := 666
	fmt.Println("a  = ", a)
}

// import "time"

func main() {

	//break// break is not in a loop ,switch, or select
	//continue//continue si not in a loop

	//goto可以用在任何地方，但是不能跨函数使用

	// fmt.Println("1111111111")

	// goto End //goto是关键字，End是用户起的名字，他叫标签

	// fmt.Println("22222222222222222")

	// End:
	// fmt.Println("3333333333333")

	MyFunc()

}
