/*反射第二定律：从反射对象到接口值的反射
Reflection goes from reflection object to inerface value

   就像物理学上的作用和反作用力，我们可以从接口值到反射对象，与此同时，我们也可以从反射对象到接口值。
   给定一个reflect.Value， 我们能用Interface方法吧他恢复成一个接口值；效果上就是这个Interface方法
把类型和值得信息打包成一个接口表示并且返回结果，简单的说，Interface就是Valueof函数的逆，除了他的返回值
的类型总是interface{}静态类型。
   重复一遍，反射就是从接口值到反射对象，然后再反射回来。*/

package main

import (
	"fmt"
	"reflect"
)

type Mytype string

func main() {

	var z Mytype = "zhengjunhua"
	Value := reflect.ValueOf(z) //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	fmt.Println(Value)

	x := Value.Interface() //我们想要的是Value里面保存的具体指，我们不需要对v.Interface方法的结果调用类型断言。
	fmt.Println(x)
}
