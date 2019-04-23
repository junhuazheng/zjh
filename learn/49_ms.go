package main

// var (
// 	size     := 1024
// 	max_size = size * 2
// )

func main() {
	println(size, max_size)
}

/*
变量简短魔石
变量简短模式限制：
1、定义变量同时显式初始化
2、不能提供数据类型
3、只能在函数内部使用
输出结果为：
syntax error : unexpected :=, expecting type
*/
/*修改后为*/
var (
	size     = 1024
	max_size = size * 2
)

//输出结果为： 1024 2048
