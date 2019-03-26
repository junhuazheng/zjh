//Go语言接口和类型之间的转换
//Go语言中使用接口断言（type assertions）将接口转换成另外一个接口，也可以将接口转换为另外的类型

//类型断言的格式：
t := i.(T)

//其中，i代表接口变量，T代表转换的目标类型，t代表转换后的变量

//如果i没有完全实现T接口的方法，这个语句将会出发宕机。出发宕机不是很友好，因此上面的语句还有一种写法：
t, ok := i.(T)

//这种写法下，如果发生接口未实现，将会把ok置为falser，t置为T类型的0值。
//正常实现ok为true，这里ok可以被认为是：i接口是否实现T类型的结果


//interface{}类型表示空接口，意思就是这种接口可以保存为任意类型

var obj interface = new(bird)
f, isFlyer := obj.(Flyer)