//包导入后的init()函数初始化顺序

package main

import (
	"chapter08/code8-2/pkg1"
	"fmt"
)

func main() {
	pkg1.ExecPkg1()
}

package pkg1

import (
	"chapter08/code8-2/pkg2"
	"fmt"
)

func ExecPkg1() {
	
	fmt.Println("ExecPkg1")
	
	pkg2.ExecPkg2()
}

func init() {
	fmt.Println("pkg1 init")
}

package pkg2

import "fmt"

func ExecPkg2() {
	fmt.Println("ExecPkg2")
}

func init() {
	fmt.Println("pkg2 init")
}
