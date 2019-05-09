//Go程序中对指针获取 反射对象 时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型。
//这个获取过程被称为取元素，等效于对指针类型变量做一个 * 操作。

package main

import (
	"fmt"
	"reflect"
)

func main() {

	type cat struct {
	}

	ins := &cat{}

	typeOfcat := reflect.TypeOf(ins)

	fmt.Printf("name:'%v' kind:'%v'", typeOfcat.Name(), typeOfcat.Kind())

	//取类型的元素
	typeOfcat = typeOfcat.Elem()

	fmt.Printf("element name'%v', element kind: '%v'\n", typeOfcat.Name(), typeOfcat.Kind())
}
