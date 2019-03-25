//Go语言实现日志系统
//日志可以用于查看和分析应用程序的运行状态。日志一半可以支持输出多种形式，如命令行、文件、网路等

//本例将搭建一个支持多种写入器的日志系统，可以自由扩展多种日志写入设备

//日志对外接口
//本例中定义一个日志写入器接口（LogWriter），要求写入设备必须遵守这个接口协议才能被日志器（Logger）注册
//日志器有一个写入器的注册方法（Logger的RegisterWriter()方法）

//日志器还有一个Log（）方法，进行日志的输出，这个函数会将日志写入到鄋已经注册的日志写入器（LogWriter）中：
package main

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writerList []LogWriter
}

//注册一个日志写入器
//注册的意思就是将日志写入器的接口添加到writerList中
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

func (l *Logger) Log(data interface{}) {

	for _, writer := range l.writerList {

		writer.Write(data)
	}
}

func NewLogger() *Logger {
	return &Logger{}
}
