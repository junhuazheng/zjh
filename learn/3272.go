//导出结构体及接口成员

//在被导出的结构体或接口中，如果他们的字段活方法首字母是大写，外部可以访问这些字段和方法：
type MyStruct struct {
	ExportedField int

	privareField int
}

type MyInterface interface {
	ExportedMethod()

	privateMethod()
}

//在代码中，MyStruct的ExportedField和MyInterface的ExportedMethod()可以被包外访问

//导入包后自定义引用的包名
package main

import (
	renameLib "chapter08/importadd/mylib"
	"fmt"
)

func main() {
	fmt.Println(renameLib.Add(1, 2))
}

//匿名导入包————只导入包但不使用包内类型和数值

//如果只希望导入包，而不使用任何包内的结构和类型，也不调用包内的任何函数时，可以使用匿名导入包：

import (
	_ "path/to/package"
)

//其中，path/to/package表示要导入的包名，下划线_表示匿名导入包
