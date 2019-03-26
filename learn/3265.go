//Go语言空接口类型（interface{}）
//空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口
//从实现的角度看，任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值

//空接口的内部实现保存了对象的类型和指针。
//使用空接口保存一个数据的过程回避直接数据对应类型的变量保存稍慢

//将值保存到空接口

var any interface{}

any = 1
fmt.Println(any)

any = "hello"
fmt.Println(any)

any = false
fmt.Println(any)

//从空接口获取值
var a int = 1

var i interface{} = a

var b int = i
//编译报错
//虽然i在复制完成后的内部值为int，单i还是一个interface{}类型的变量
//使用类型断言修改：
var b int = i.(int)
//修改后，代码可以编译通过，并且b可以获得i变量保存的a变量的值：1