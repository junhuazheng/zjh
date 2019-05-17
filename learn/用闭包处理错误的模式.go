package main

import "fmt"

func A() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in A", r)
		}
	}()

	fmt.Println("Calling A")
	B(0)
	fmt.Println("Returned normally form g")

}

func B(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in B", i)
	fmt.Println("Printing in B", i)
	B(i + 1)
}

func main() {
	A()
	fmt.Println("程序结束!")
}
