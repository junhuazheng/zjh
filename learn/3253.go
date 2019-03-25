//命令行写入器
//在UNIX思想中，一切皆为文件。文件包括内存、磁盘、网络和命令行等。这种抽象方法方便我们访问这些看不见摸不着的虚拟资源
//命令行在GO中也是一种文件，os.Stdout对应标准输出，一般表示屏幕有，也就是命令行
//os.Stderr对应错误输出，一般将错误输出到日志中。
//os,Stdin对应标准输出，一般表示键盘。
//os.Stdout、os.Stderr、os.Stdin都是*os.File类型，和文件一样实现了io.Writer接口和Write()方法

//将命令行抽象为日志写入器
package main

import (
	"fmt"
	"os"
)

type consoleWriter struct {
}

func (f *consoleWriter) Write(data interface{}) error {

	str := fmt.Sprintf("%v\n", data)

	//与fileWriter类似，这里也将str字符串转换为字节数组并写入标准输出os.Stdout
	//写入后的内容就会显示在命令行中
	_, err := os.Stdout.Write([]byte(str))

	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
