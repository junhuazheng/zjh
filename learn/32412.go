package main

//在代码中使用Socket结构实现Writer接口和Closer接口代码如下：

//使用io.Writer的代码，并不知道Socket和io.Closer的存在
func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

//使用io.Closer，并不知道Socket和io.Writer的存在
func usingCloser(closer io.Closer) {
	closer.Close()
}

func main() {
	s := new(Socket)

	usingWriter(s)

	usingCloser(s)
}

//usingWriter()和usingCloser()完全独立，互相不知道对方的存在，也不知道自己使用的借口是Socket实现的
