/*
空interface不包含任何的method，正因为如此，所有的类型都实现了空interface。
空interface对于描述起不到任何的作用，但是空interface在我们需要存储任意类型
的数值的时候相当有用，因为它可以存储任意类型的数值，一个函数吧interface{}作为
参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{}，那么也就
可以返回任意类型的值。
*/

package main

import (
	"fmt"
	"reflect"
)

func Myecho(a interface{}) {
	fmt.Printf("变量的值是: \033[31;1m[%v]\033[0m, 其类型是:\033[31;1m[%v]\033[0m\n", a, reflect.TypeOf(a))

}

func main() {
	Name := "郑俊华"
	Age := 25
	var b string
	b = "Golang"
	Language := []string{b}
	fmt.Println(reflect.TypeOf(Name), reflect.TypeOf(Age), reflect.TypeOf(Language))

	var zhengjunhua interface{}

	zhengjunhua = Name
	Myecho(zhengjunhua)

	zhengjunhua = Age
	Myecho(zhengjunhua)

	zhengjunhua = Language
	Myecho(zhengjunhua)
}
