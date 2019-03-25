//文件写入器
//文件写入器（fileWriter）是众多日志写入器（LogWriter）中的一种。
//文件写入器的功能是根据一个文件名创建日志文件（fileWriter的SetFile方法）。
//在有日志写入时，将日志写入文件中
package main

import (
	"errors"
	"fmt"
	"os"
)

//声明文件写入器，在结构体中保存一个文件句柄，方便每次写入时操作
type fileWriter struct {
	file *os.File
}

//设置文件写入器写入的文件名
func (f *fileWriter) SetFile(filename string) (err error) {

	//如果文件已经打开，关闭前一个文件
	//考虑到SetFile()方法可以被多次调用（函数可重入性），假设之前已经调用过SetFile()
	//后再次调用，此时的f.flie不为空，就需要关闭之前的文件，重新创建新的文件
	if f.file != nil {
		f.file.Close()
	}

	//创建一个文件并保存文件句柄
	f.file, err = os.Create(filename)

	//如果创建的过程出现错误，则返回错误
	return err
}

//实现LogWriter的Write（）方法
func (f *fileWriter) Write(data interface{}) error {

	//如果文件没有准备好，文件句柄为nil，此时使用errors包的NEw()函数返回一整个错误对象，包含一个字符串"file not created"
	if f.file == nil {

		//日志文件没有准备好
		return errors.New("file not created")
	}

	//使用fmt.Sprintf将data转换为字符串，这里使用的格式化参数是%v，意思是讲打他按期本来的值转换为字符串
	str := fmt.Sprintf("%v\n", data)

	//将数据以字节数组写入文件中
	_, err := f.file.Write([]byte(str))

	return err
}

//创建文件写入器实例
func newFileWriter() *fileWriter {
	return &fileWriter{}
}
