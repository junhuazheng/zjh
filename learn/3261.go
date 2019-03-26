//Go语言接口的嵌套组合
//Go语言中，接口与接口间也可以通过嵌套创造出新的接口
//接口与接口嵌套组合而成的新接口，只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可被调用

//Go语言的io包中定义了写入器（Writer）、关闭器（Closer）和写入器（WriteCloser）3个接口
type Writer interface {
	Write(p []byte) (n int, err error)
}