//使用 reflect.TypeOf() 函数可以获得任意值的类型对象(reflect.Type)
//程序通过类型类型对象可以访问任意值的类型信息。
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int

	typeOfA := reflect.TypeOf(a)

	fmt.Println(typeOfA.Name(), typeOfA.Kind())

}
