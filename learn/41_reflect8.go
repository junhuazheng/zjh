package main

import (
	"fmt"
	"reflect"
)

type Gamer struct {
	Num  int
	Name string
	Lv   int
}

func (g Gamer) Fangfa1(name string, lv int) {
	fmt.Println("Gamer name: ", name, "lv: ", lv, "origal Gamer.Name: ", g.Name)
}

func (g Gamer) Fangfa2() {
	fmt.Println("Fangfa2")
}

func main() {
	gamer := Gamer{99778, "OldMa", 100}

	getValue := reflect.ValueOf(gamer)

	methodValue := getValue.MethodByName("Fangfa1")
	ff := []reflect.Value{reflect.ValueOf("SamllMa"), reflect.ValueOf(99)}
	methodValue.Call(ff)

	methodValue = getValue.MethodByName("Fangfa2")
	ff = make([]reflect.Value, 0)
	methodValue.Call(ff)
}

/*
1、要通过反射来调用其对应的方法，必须要先通过reflect.Value(interface)
来获取到reflect.Value得到“反射类型对象”后才能做下一步处理
2、reflect.Value.MethodByName，需要指定准确真实的方法名字，如果错误将直接panic
MethodByName返回一个函数值对应的reflect.Value方法的名字。
3、[]reflect.Value，这个是最终需要调用的方法的参数，可以没有或者一个或者多个，根据实际参数来定
4、reflect.Value的 Call 这个方法，这个方法将最终调用真实的方法，参数请保持一致
如果reflect.Value.kind是一个方法，那么将直接panic
5、本来可以用u.FangfaX直接调用的，但是如果要通过反射，那么首先要将方法注册
也就是MethodByName，然后通过反射调	methodValue.Call。


射可以大大提高程序的灵活性，使得interface{}有更大的发挥余地
反射必须结合interface才玩得转
变量的type要是concrete type的（也就是interface变量）才有反射一说
反射可以将“接口类型变量”转换为“反射类型对象”
反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
反射可以将“反射类型对象”转换为“接口类型变量
reflect.value.Interface().(已知的类型)
遍历reflect.Type的Field获取其Field
反射可以修改反射类型对象，但是其值必须是“addressable”
想要利用反射修改对象状态，前提是 interface.data 是 settable,即 pointer-interface
通过反射可以“动态”调用方法
因为Golang本身不支持模板，因此在以往需要使用模板的场景下往往就需要使用反射(reflect)来实现
*/
