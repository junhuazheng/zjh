//Go语言的源码复用建立在包基础上。
//Go语言的入口main()函数所在的包（package）叫main，main包想要引用别的代码，必须同样以包的方式引用

//Go语言的包与文件已一一对应，所有与包相关的操作，必须依赖于工作目录（GOPATH）

//导入包
// package main

// import (
// 	"chapter08/importadd/mylib"
// 	"fmt"
// )

// func main() {
// 	fmt.Println(mylib.Add(1, 2))
// }

//在Go语言中，如果想在一个包里引用另一个包的标识符（如类型、变量、常亮等）时，必须首先将被引用的标识符到处
//将要导出的标识符的首字母大写就可以让引用者可以访问这些标识符了

//导出包内标识符
//下面代码中包含一系列未导出标识符，他们的首字母都是小写，这些标识符可以在包内自由使用，但是包外无法访问他们
package mypkg

var myVar = 100

const myConst = "hello"

type myStruct struct {
}

//将myStruct和myConst首字母大写，导出这些标识符

package mypkg

var myVar = 100

const MyConst = "hello"

type MyStruct struct {
	
}

//此时，MyConst和MyStruct可以被外部访问，而myVar犹豫首字母是小写，因此只能在mypkg包内使用，不能被外部包引用
