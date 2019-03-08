package main

import "fmt"

func main() {
	sum := 1
	for i := 2; i <= 200; i++ {
		sum = sum + i
	}
	fmt.Println("sum = ", sum)

	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	}

	b := 10
	if b == 10 {
		fmt.Println("b == 10")
	}
	if b > 10 {
		fmt.Println("b > 10")
	}
	if b < 10 {
		fmt.Println("b < 10")
	}

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

	var num int
	//num := 1
	fmt.Printf("请按下楼层: ")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("按下的是1楼")
		fallthrough
	case 2:
		fmt.Println("按下的是2楼")
		//break
		fallthrough
	case 3:
		fmt.Println("按下的是3楼")
		//break
		fallthrough
	case 4:
		fmt.Println("按下的是4楼")
		//break
		fallthrough
	default:
		fmt.Println("按下的是xxx楼")
	}
}
