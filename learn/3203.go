//匿名函数视线操作封装
//下面这段代码将匿名函数作为map的键值，通过命令行参数动态调用匿名函数：
package main

import (
	"flag"
	"fmt"
)

//定义命令行参数skill，从命令行输入-skill可以将空格吼得字符串传入skillParam指针变量
var skillParam = flag.String("skill", "", "skill to perform")

func main() {

	//解析命令行参数，解析完成后，skillParam指针变量将指向命令行传入的值
	flag.Parse()

	//定义一个字符串映射到func()的map，然后填充这个map
	var skill = map[string]func(){
		//初始化map的键值对，值为匿名函数
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("angel fly")
		},
		"fly": func() {
			fmt.Println("soldier run")
		},
	}

	//skillParam是一个*string类型的指针变量，使用*skillParam获取到命令行传过来的值
	//并在map中查找对应命令行参数指定的字符串的函数。
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found") //如果在map定义中存在这个参数就调用；否则就打印“技能没有找到“
	}
}
