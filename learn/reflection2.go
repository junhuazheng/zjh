package main

import (
	"fmt"
	"reflect"
)

type MyType string

func main() {
	var y MyType = "zhengjunhua"
	Type := reflect.TypeOf(y)
	Value := reflect.ValueOf(y)

	fmt.Println("type\t\t\t:", Type)
	fmt.Println("underlying type :", Type.Kind())

	fmt.Println("value\t\t\t:", Value)
	fmt.Println("static type :", Value.Type()) //Value类型有一个Type方法可以返回reflect.Value类型的Type（这个返回的是静态类型即“static type”）

	fmt.Println("underlying type:", Value.Kind())
	fmt.Println("kind is string:", Value.Kind() == reflect.String)

	fmt.Println("value\t\t\t:", Value.String()) //通过Value类型String方法来让我们获取存在里面的值。如果底层数据是“int”就用“Int”方法
}
