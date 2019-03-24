//类型和接口之间有一对多和多对一的关系

//一个类型可以实现多个接口
//一个类型可以同事实现多个接口，而接口间彼此独立，不知道对方的实现

//网络上的两个程序通过一个双向的通信连接实现数据的交换，连接的一端称为一个Socket。
//Socket能够同时读取和写入数据，这个特性与文件类似。因此开发中把文件和Soket都具备的读写特性抽象为独立了的读写器概念

//Socket和文件一样，在使用完毕后，也需要对资源进行释放。
//把Socket能够写入数据和需要关闭的特性使用接口来描述：
type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (s *Socket) Close() error {
	return nil
}

//Socket结构的Write()方法实现了io.Writer接口：
type Writer interface {
	Write(p []byte) (n int, err error)
}

//同时，Socket结构也实现了io.Closer接口：
type Closer interface {
	Close() error
}

//使用Socket实现的Writer接口的代码，无须了解Writer接口的实现者是否具备Closer接口的特性
//同样，使用Closer接口的代码也并不知道Socket已经实现了Writer接口

