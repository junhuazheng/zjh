/*现在一些流星设计思想需要建立在反射基础上，如控制翻转（Inversion Of Control， IOC）
依赖注入（Dependency Injection， DI）
Go语言中非常有名的Web框架martini就是通过依赖注入技术进行中间件的视线
使用martni框架塔尖的http的服务器：*/
package main

import "github.com/go-martinni/martini"

func main() {

	m := martini.Classic()

	//响应路径/的代码使用了一个闭包实现。
	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Run()
}

/*如果希望蝴蝶Go语言提供的请求和响应接口，可以直接修改为：
m.Get("/", func(res http.ResponseWriter, req *http.Request) string {
	响应处理代码......
	})

martini的底层会自动通过识别Get获得的闭包参数情况，通过动态发射调用这个函数病传入需要的参数。
martini的设计广受好评，但其运行效率较低。其中最主要的因素是大量使用了反射。

虽然一般情况下，I/O的延迟远远大于反射代码所造成的延迟。
但是，更低的响应速度和更低的CPU占用依然是Web服务器追求的目标。
因此，反射在带来灵活性的同事，也戴上了性能低下的桎梏。*/
