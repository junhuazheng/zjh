//通过反射获取类型信息
//使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type）
//程序通过类型对象可以访问任意值的类型信息

package main

import (
	"fmt"
	"reflect"
)

type Enu int

const (
	Zero Enu = 0
)

func main() {

	var a int

	typeOA := reflect.TypeOf(a)

	//通过typeOfA类型对象的成员函数，可以分别获取到typeOfa变量的类型名和种类
	fmt.Println(typeOA.Name(), typeOA.Kind())

	type cat struct {
	}

	typeOfCat := reflect.TypeOf(cat{})

	fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) //输出结果: cat struct

	typeOfA := reflect.TypeOf(Zero)

	fmt.Println(typeOfA.Name(), typeOfA.Kind()) //输出结果: Enu int
}

//Go语言中的类型（Type）指的是系统原生数据类型，以及使用type关键字定义的类型， 这些类型的名称就是其类型本身的名称
//种类（Kind）指的是对象归属的品种
//reflect.Type中的Name()方法，返回表示类型名称的字符串
//reflect.Type中的Kind()方法，返回reflect.Kind类型的常量
