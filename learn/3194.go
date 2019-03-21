package main

//字符串处理的主要流程包含以下几个步骤：
//1、准备要处理的字符串列表
//2、准备字符串处理链
//3、处理字符串列表
//4、打印输出后的字符串列表

func main() {

	//待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	//处理函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace, //移除字符串前缀左边的一个空格，这个函数的定义行号符合处理函数的格式
		Strings.ToUpper,   //用于将字符串转为大写
	}

	//处理字符串
	StringsProccess(list, chain)

	//输出处理好的字符串
	for _, sht := range list {
		fmt.Println(str)
	}
}
