//类型别名作为结构体嵌入的成员会发生什么情况
package main

import (
	"fmt"
	"reflect"
)

type Brand struct {
}

func (t Brand) Show() {

}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {

	var a Vehicle

	a.FakeBrand.Show()

	ta := reflect.TypeOf(a) //使用反射取变量a的反射类型对象，以查看其成员类型。

	for i := 0; i < ta.NumField(); i++ {

		f := ta.Field(i)

		fmt.Printf("FildName: %v, FildType: %v\n", f.Name, f.Type.Name())
	}
}

//这个例子中，FakeBrand是Brand的一个别名。在Ve中嵌入FakeBrand和Brand并不意味着嵌入两个Brand。
//FakeBrand的类型会以名字的方式保留在Ve的成员中。
