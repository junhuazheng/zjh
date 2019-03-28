//Go语言package（创建包）

//包（package）是多个Go源码的集合，是一种高级的代码复用方案
//Go语言默认为我们提供了很多包，如fmt、os、io包等，开发者可以根据自己的需要创建自己的包

//包要求在同一个目录下的所有文件的第一行添加代码，以标记该文件归属的包：
package 包名

//包的特性：
//一个目录下的同级稳健归属一个包
//包名可以与其目录不同名
//包名为main的包为应用程序的入口包，编译源码没有main包是，将无法编译输出可执行文件

//在被导出的结构体或接口中，如果他们的字段活方法首字母是大写，外部可以访问这些字段和方法：

type MYStruct struct {

	//包外可以访问的字段
	ExporyedField int

	//仅限包内访问的字段
	privateField int
}

type MyInterface interface {

	//包外可以访问的方法
	ExportedMethod()

	//仅限包内访问的方法
	privateMethod()
}
