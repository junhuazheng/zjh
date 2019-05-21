//反射第三定律：要修改一个反射对象，该值必须是可设置的（settable）
//To modify a reflection object,the value must be settable

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var zjh string = "zhengjunhua"
	p := reflect.ValueOf(&zjh)                   //这里将zjh地址传进去了
	fmt.Println("type of p :", p.Type())         //我们传进去的是地址，所以得到的应该是一个指针类型的string
	fmt.Println("settability of p:", p.CanSet()) //反射对象p不是settable的，因此返回值应该是一个false

	v := p.Elem()
	//反射对象p不是settable的，但是我们想要设置的不是p，而是（效果上来说）*p
	//为了得到p指向的东西，我们调用Value的Elem方法。

	fmt.Println(v.Interface()) //查看v里面的值
	s := v.String()
	s = "郑俊华" //我们此处修改的只是“zjh”变量中的一个副本
	fmt.Println(s)
	fmt.Println(zjh)

	fmt.Println("settability of v:", v.CanSet()) //反射对象v是settable的，因此返回值应该是一个true
	v.SetString("Golang")                        //调用SetString,SetInt,SetFloat等方法去修改源数据
	fmt.Println(zjh)
}
