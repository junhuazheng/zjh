//Go语言中的类型名称对应的反射获取方法是 reflect.Type 中的 Name() 方法，返回表示类型名称的字符串。
//类型归属的种类(Kind)使用的是 reflect.Type 中的 Kind() 方法，返回reflect.Kind 类型的变量。

package main

import (
	"fmt"
	"reflect"
)

type Enum int

const (
	Zero Enum = 0
)

func main() {

	type cat struct {
	}

	//获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})

	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	typeOfA := reflect.TypeOf(Zero)

	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
