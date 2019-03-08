package main

import "fmt"

func main() {
	s := "猪"

	if s == "猪" {
		fmt.Println("从幼稚到成熟")
	}
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}

	sum := 10

	for i := 11; i <= 50; i++ {
		sum = sum + i
	}
	fmt.Println("sum = ", sum)

	str := "zjh"

	for j := 0; j < len(str); j++ {
		fmt.Printf("str[%d] = %c\n", j, str[j])
	}
	for j, data := range str {
		fmt.Printf("str[%d]=%c\n", j, data)
	}

	sum1 := 1
	for k := 1; k <= 10; k++ {
		sum1 = sum1 * k
	}
	fmt.Println("sum1 = ", sum1)
}
