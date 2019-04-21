/*
从reflect.Value中获取借口interface的信息
当执行reflect.ValueOf(interface)之后,就得到了一个类型为"reflect.Value"变量
可以通过它本身的Interface()方法获得接口变量的真实内容，然后可以通过类型判断进行转换
转换为原有的真实类型。
不过，我们可能是已知原有类型，也有可能是未知原有类型，因此下面分两种情况进行说明。

已知原有类型进行“强制转换”
已知类型后转换为其对应的类型的做法如下，直接通过Interface方法然后强制转换：
realValue := value.Interface() (已知类型)
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.222

	pointer := reflect.ValueOf(&num) /*反射类型对象*/
	value := reflect.ValueOf(num)

	//可以理解为“强制转换”，但是需要注意的是，转换的时候，如果转换的类型不完全符合，则直接panic
	//Golang对类型要求非常严格，类型一定要完全符合

	//如下两个，一个是*float64，一个是float64，如果弄混则会panic
	conP := pointer.Interface().(*float64) /*接口类型变量*/
	conV := value.Interface().(float64)

	fmt.Println(conP)
	fmt.Println(conV)
}

/*
1、转换时，如果转换的类型不完全符合，则直接panic
2、也就是说，反射可以将“反射类型对象”重新转换为“接口类型变量”
*/
