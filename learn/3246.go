//开发中常用的接口及写法
//Go语言提供的很多包中都有接口，例如io包中提供的Writer接口：
type Writer interface {
	Write(p []byte) (n int, err error)
}

//这个接口可以调用Write()方法写入一个字节数组（[]byte），返回值告知写入字节数（n int）和可能发生的错误（err error）

//类似的们还有将一个对象以字符串形式展现的接口，只要实现了这个接口的类型
//在调用String()方法时，都可以获得对象对应的字符串。在fmt包中的定义如下：
type Stringer interface {
	string() string
}

//Stringer接口在Go语言中的使用频率非常高
//Go语言的每个接口中的方法数量不会很多。
//GO语言希望通过一个接口精准描述它自己的功能
//而通过多个接口的嵌入和组合的方式将简单的借口扩展为复杂的借口