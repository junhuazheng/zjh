//使用指针
//1、定义一个指针变量
//2、将一个变量的地址赋值给一个指针
//3、最后访问指针变量中可用地址的值
package main

import "fmt"

func main() {

	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("Address of a variable: %x\n", &a)

	fmt.Printf("%x\n", ip)

	fmt.Printf("%d\n", *ip)

	//Go编译器为指针变量分配一个Nil值，以防指针没有确切的地址分配。
	//变量声明完成时候，制定为nil值的指针为nil指针
	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr) //输出结果为0

	/*在大多数操作系统上，程序允许访问地址0处的内存，因为该内存是由操作系统保留的
	然而，存储器地址0具有特殊意义：
	它表示指针不打算指向可访问的存储器位置。但是按照惯例，如果指针包含nil值，则假设它不指向任何东西*/

	var b int = 1

	var boy *int
	boy = &b

	if boy != nil {
		fmt.Printf("%x\n", *boy)
		return
	}

}
