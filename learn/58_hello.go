package main

import "fmt"

func MyFunc01(a int) {
	//a = 111
	fmt.Println("a = ", a)
}

func MyFunc02(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc03(a, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc04(a int, b string, c float64) {

}

func MyFunc05(a, b string, c float64, d, e int) {

}

func MyFunc06(a string, b string, c float64, d int, e int) {

}

func main() {
	MyFunc01(666)

	MyFunc02(666, 777)
}
