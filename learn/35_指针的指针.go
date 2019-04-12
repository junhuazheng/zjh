/*指向指针的指针是多重间接的形式或指针链。
通常，指针包含变量的地址。当定义指向指针的指针时：
第一个指针包含第二个指针的地址，它指向包含实际值的位置。

作为指向指针的指针的变量必须如此声明，
这是通过在其名称前添加一个星号（*）来实现的。
例如，一下是一个指向int类型的指针的声明：
var ptr **int
当目标值由指向指针的指针间接指向时，访问该值需要应用两个星号（**）运算符：*/
package main

import "fmt"

func main() {

	var a int

	var ptr *int

	var pptr **int

	a = 3000

	ptr = &a

	pptr = &ptr

	fmt.Printf("value of a = %d\n", a)

	fmt.Printf("Value available at *ptr = %d\n", *ptr)

	fmt.Printf("Value available at **pptr = %d\n", **pptr)
}
