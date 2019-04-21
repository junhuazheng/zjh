/*
reflect的基本功能TypeOf和ValueOf
既然反射就是用来检测存储在借口变量内部（值value类型concrete type）pair对的一种机制。
那么在Golang的reflect范社保中有什么那个的方式可以让我们直接获取到变量内部的信息呢？
它提供了两种类型（或者说两个方法）让我们可以很容易的访问接口变量内容，分别是reflect.ValueOf()和relflect.TypeOf().

func ValueOf(i interface{}) Value{...}
ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0

func TypeOf(i interface{}) Type {...}
TypeOf用来动态获取输入参数接口中的值类型，如果接口为空则返回nil

(都是用来获取输入参数接口中的数据）（pair的值和类型）
reflect.TypeOf()是获取pair中的type，reflect.ValueOf()获取pair中的value
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.222

	fmt.Println("type: ", reflect.TypeOf(num)) //这里直接使用获取pair的值

	fmt.Println("value: ", reflect.ValueOf(num))

}

/*
运行结果为：
type : float64
value : 1.222

1、reflect.TypeOf:直接给到了我们想要的type类型，如float64、int、各种pointer、struct等等真实类型。
2、reflect.ValueOf:直接给到了我们想要的具体值，如1.222这个具体数值，或者类似&{1"Allen.Wu"25}这样的结构体struct的值。
3、也就是说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种。
*/
