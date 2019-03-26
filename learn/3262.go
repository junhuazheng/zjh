//io.WriteCloser接口在使用时，编译器会根据接口的实现者确定他们是否同时实现了io.Writer和io.Closer接口
package main

import "io"

type device struct {
}

func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func main() {

	//对device实例化，犹豫device实现了io.WriteCloser的所有嵌入接口
	//因此device指针就会被隐式转换为io.WriteCloser接口
	var wc io.WriteCloser = new(device)

	wc.Write(nil)

	wc.Close()

	var writeOnly io.Writer = new(device)

	writeOnly.Write(nil)
}
