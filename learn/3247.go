//Go语言实现接口的条件
//接口定义吼，需要实现接口，调用方才能正确编译通过并使用接口
//接口的实现需要遵循两条规则才能让接口可以

//1、接口的方法与实现接口的类型方法格式一致
//在类型中添加与接口签名一致的方法可以实现该方法。
//签名包括方法中的名称、参数列表、返回参数列表。
//也就是说，只要实现接口类型中的方法的名称、参数列表、返回参数列表中的任意一项与接口要实现的方法不一致，那么这个接口的方法就不会被实现

//为了抽象数据写入的过程，定义DataWriter接口来描述数据写入需要实现的方法
//接口中的WriteData()方法表示将数据写入，写入方无须关心写入到那里。
//实现接口的类型实现WriteData方法时，会具体编写将数据写入到什么结构中。
//这里使用file结构体实现DataWriter接口的WriteData方法，方法内部只是打印一个日志，表示有数据写入

//数据写入器的抽象：
package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {

	fmt.Println("WriteData:", data)
	return nil
}

func main() {
	f := new(file)

	var writer DataWriter

	writer = f

	writer.WriteData("data")
}
