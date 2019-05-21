//反射第一定律：从接口值到发射对象的反射
//Reflection goes from interface value to reflection object

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var zjh float64 = 13.8

	fmt.Println("type: ", reflect.TypeOf(zjh))
	/*reflect.Typeof 签名里面就包含了一个空接口。当我们调用reflect.Typeof(zjh)的时候
	zjh首先被保存到一个空姐口中，这个空接口被作为参数传递。reflect.Typeof会把这个空接口拆包（unpack）恢复出类型信息*/

	fmt.Println("value:", reflect.ValueOf(zjh))
	//reflect.Valueof可以把值恢复出来，Valueof方法会返回一个Value类型的对象对象。
}

/*reflect.Type和reflect.Value这两种类型都提供了大量的方法让我们可以检查和操作这两种类型，需要注意两点：

1、Value类型有一个TyPe方法可以返回reflect.Value类型的Type，这个方法返回的是值得静态类型
即“static type”，也就是说如果定义了“type MyType string”那么这个函数返回的是MyType类型而不是string。
有关Value类型的带有名字注入“String”，“Int”，“Uint”“Bytes”等等的方法可以让我们获取存在里面的值。

2、Type和Value都有一个Kind方法可以返回一个常量用于指示一个项到底是以什么形式存储的，也
就是底层类型“underlying type”。这些常量包括：Unit，Float64,Slice等等。

具体用法看2*/
