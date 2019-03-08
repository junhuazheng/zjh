package main

import "fmt"

func main() {

	//for 初始化条件； 判断条件 ； 条件变化{
	//}
	//1+2+3 ...... 100累加

	sum := 0

	//1）初始化条件 i := 1
	//2）判断条件是否为真， i <=100, 如果为真，执行循环体，如果为假，跳出循环
	//3）条件变化 i++
	//4)重复2， 3， 4
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println("sum = ", sum)

	str := "abc"

	//通过for打印每一个字符
	for j := 0; j < len(str); j++ {
		fmt.Printf("str[%d] = %c\n", j, str[j])
	}

	//迭代打印每一个元素，默认返回2个值：一个是元素的位置，一个是元素本身
	for j, data := range str {
		fmt.Printf("str[%d]=%c\n", j, data)
	}

	for j := range str { //第2个返回值，默认丢弃，返回元素的位置（下标）
		fmt.Printf("str[%d]=%c\n", j, str[j])
	}

	for j, _ := range str { //第2个返回值，默认丢弃，返回元素的位置（下标）
		fmt.Printf("str[%d]=%c\n", j, str[j])
	}
}
