//使用日志
//在程序中使用日志器一般会先通过代码创建日志器（Logger），为日志器添加输出设备
//fileWriter、consoleWriter等
//这些设备有一部分需要一些参数设定，如文件日志写入器需要提供文件名（SetFile()）

//使用日志器：
package main

import "fmt"

func createLogger() *Logger {

	l := NewLogger()

	cw := newConsoleWriter()

	l.RegisterWriter(cw)

	fw := newFileWriter()

	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}

	l.RegisterWriter(fw)

	return l

}

func main() {
	l := createLogger()

	l.Log("hello")
}
