匿名函数定义格式：
func(参数列表)(返回参数列表){
	函数体
}

func(data int) {
	fmt.Println("hello", data)
}(100)

f := func(data int) {
	fmt.Println("hello", data)
}

f(100)