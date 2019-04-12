package main

/*递归是以自相似的方式重复项的过程。
在编程语言中允许在函数内调用同一个函数称为递归调用。
func recursion() {
	recursion()
}

func main() {
	recursion()
}
*/
/*Go编程语言支持递归，即函数调用自身的函数。
但是在使用递归时，程序员㤇注意在函数中定义或设置一个退出条件，否则他会进入无限循环！
递归函数非常有用，可用于解决许多数学问题，如计算数学的竭诚，生成斐波那契数列等。*/

//数字阶乘
import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {
	var i int = 15
	fmt.Printf("Factorial of %d is %d", i, factorial(i))
}
