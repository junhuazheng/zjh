package main

import (
	"log"
	"os"
)

/*你也可以自定义Logger类型，log.Logger提供了一个New方法用来创建对象

func New(out io.Writer, prefix string, flag int) *Logger

该函数共有三个参数：
1、输出位置out，是一个io.Writer对象，该对象可以使一个文件也可以是实现了该接口的对象。
通常我们可以用这个来指定日志输出到哪个稳健。
2、prefix我们在前面已经看到，就是日志内容前面的东西。我们可以将其置为"[Info]"、"[Warning]"等来帮助区分日志级别
3、flags是一个选项，显示日志开头的东西，可选值有：

Ldate         = 1 << iota     // 形如 2009/01/23 的日期
Ltime                         // 形如 01:23:23   的时间
Lmicroseconds                 // 形如 01:23:23.123123   的时间
Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23
Lshortfile                    // 文件名和行号: d.go:23
LstdFlags     = Ldate | Ltime // 日期和时间*/

func main() {
	fileName := "Info_First.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	debuglog := log.New(logFile, "[Info]", log.Llongfile)
	debuglog.Println("A Info message here")
	debuglog.SetPrefix("[Debug]")
	debuglog.Println("A Debug Message here")
}
