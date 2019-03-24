package main

import "fmt"

//定义DataWriter接口。这个接口只有一个方法，即WriteData()
//输入一个interface{}类型的data，返回一个error结构表示可能发生的错误
type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

//file的WriteData()方法使用指针接收器。输入一个interface{}类型的data，返回error
func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}

func main() {
	//实例化file赋值给f，f的类型为*file
	f := new(file)

	//声明DataWriter类型的writer接口变量
	var writer DataWriter

	//将*file类型的f赋值给DataWriter接口的writer，虽然两个变量类型不一致
	//但是writer是一个接口，且f已经完全实现了DataWriter()所有方法，因此赋值是成功的
	writer = f

	//DataWriter接口类型的writer使用WriteData()方法写入一个字符串
	writer.WriteData("data")
}

//使用file结构体实现DataWriter接口的WriteData方法，方法内部只是打印一个日志，表示有数据写入
//(d *file) WriteData(data interface{}) error  实现DataWriter的WriteData()方法
