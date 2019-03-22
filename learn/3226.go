//req.Header的类型为http.Header，就是典型的自定义类型，并且拥有自己的方法。http.Header的部分定义：

//Header实际是一个以字符串为键、字符串切片为值得映射
type Herade map[string][]string

//Add()方法为Header的方法，map是一个引用类型，因此即便使用（h Header）非指针接收器，也可以修改map值
func (h Header) Add(key, value string) {
	textproto.MiMEheader(h).Add(key, value)
}

func (h Header) Set(key, value string) {
	textproto.MIMEHeader(h).Set(key, value)
}

func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}

//为类型添加方法的过程是一个语言层特性，使用类型方法的代码经过编译器编译后的代码运行效率
//与传统的面向过程或面向对象的代码没有任何区别。因此，为了代码便于理解，可以在编码时使用Go语言的类型方法特性