//准备一个字符串
str := "hell world"

//创建一个匿名函数
foo := func() {

	//匿名函数中访问str
	str = "hello dude"
}

//调用匿名安函数
foo()

