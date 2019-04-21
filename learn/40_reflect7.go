/*通过reflect.Valueof来进行方法的调用
这算是一个高级用法了，前面我们只说到对类型、变量的几种反射的用法
包括如何获取其值、其类型。如何重新设置新值。
单是在工程应用中，另外一个常用并且属于高级的用法，就是通过reflect来进行方法（函数）的调用。
比如我们要做框架工程的时候，需要可以随意扩展方法，或者说用户可以自定义方法，那么我们通过什么手段来扩展让用户能够自定义呢？
关键点在于用户的自定义方法是未知的，因此我们可以通过reflect来搞定。
*/

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Fangfa1(name string, age int) {
	fmt.Println("fangfa name: ", name, "age: ", age, "and origal User.Name: ", u.Name)
}

func (u User) Fangfa2() {
	fmt.Println("Fangfa1")
}

/*如何通过反射来进行方法的调用？
本来可以用u.FangfaX直接调用的，但是如果要通过反射
那么首先要将方法注册，也就是MethodByName
然后通过反射调动mv.Call*/

func main() {
	user := User{1, "Allen.wu", 25}

	/*要用过反射来调用其对应的方法，必须要先通过reflect.ValueOf(inertface)
	来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理。*/
	getValue := reflect.ValueOf(user)

	//1、一定要指定参数为正确的方法名
	//2、先看看带有参数的调用方法
	methodValue := getValue.MethodByName("Fangfa1")
	ff := []reflect.Value{reflect.ValueOf("wukong"), reflect.ValueOf(999)}
	methodValue.Call(ff) //调用Fangfa1()
	methodValue = getValue.MethodByName("Fangfa2")
	ff = make([]reflect.Value, 0)
	methodValue.Call(ff) //调用Fangfa2()
	//methodValue.Call want []reflect.Value 需要“反射类型对象”
}
