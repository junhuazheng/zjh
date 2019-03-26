//Go语言的switch不仅可以像其他语言一样实现数值、字符串的判断，还有一种特殊的通途————判断一个接口内保存或实现的类型

//类型断言的书写格式
//switch实现类型分支时的写法格式如下：
// switch 接口变量.(type) {
// 	case 类型1:
// 	//变量是类型1时的处理
// 	case 类型2:
// 	//变量是类型2时的处理
// 	...
// 	default:
// 	//变量不是所有case中列举的类型时的处理
// }

//接口变量：表示需要判断的接口类型的变量
//类型1：表示接口变量可能具有的类型别聊，满足时，会指定case对应的分支进行吃力

package main

import "fmt"

func printType(v interface{}) {

	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

func main() {
	printType(1024)
	printType("pig")
	printType(true)
}
