/*
通过reflect.Value设置实际变量的值
reflect.Value是通过reflect.ValueOf(X)获得的，只有当X是指针的时候
才可以通过reflect.Value修改实际变量X的值，即：
要修改反射类型的对象就一定要保证其值是“addressable“的*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num float64 = 1.2345
	fmt.Println("old value of pointer :", num)

	//通过reflect.ValueOf来获取num中的reflect.Value，注意，参数必须是指针才能修改其值。
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet()) //查询是否可以设置

	//重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)

	/*如果reflect.ValueOf的参数不是指针，即:
	pointer := reflect.ValueOf(num)
	这里会直接pannic，"panic:reflect:call of reflect.Value.Elem on float64 Value"*/
}

/*
1、需要传入的参数是*float64这个指针，然后可以通过pointer.ELem()去获取所指向的Value
2、如果传入的参数不是指针，而是变量
1）通过Elem获取原始值对应的对象则直接panic
2）通过CanSet方法查询是否可以设置返回false
3、newValue.CanSet()表示是否可以重新设置其值，输出true则可修改，输出false表示不能，修改完之后再进行打印发现真的已经修改了。
4、reflect.Value.Elem()表示获取原始值对应的反射对象，只有原始对象才能修改，当前发射对象不能修改。
5、也就是说如果要修改反射类型独享，其值必须是“addressable”
（对应的要传入的是指针，同时要通过Elem方法获取原始值对应的反射对象）
6、struct或者struct的嵌套一样的判断处理
*/
