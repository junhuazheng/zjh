//多个类型可以实现相同的接口
//一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
//也就是说，使用者不关心某个接口的方法是通过一个类型完全实现的，还是通过多个结构嵌入到一个结构体中拼凑起来共同实现的

//Service接口定义了两个方法：
//一个是开启服务的方法（Star()），一个是输出日志的方法（Log()）
//使用GameServic结构体来实现Service，GameService自己的结构只能实现Strat()方法
//而Service接口中的Log()方法已经被一个能输出日志的日志器(Logger)实现了
//无须再进行GameService封装，或者重新实现一边。
//所以，选择将Logger嵌入到GameService能最大程度避免代码冗余，简化代码结构
type Service interface {
	Start()
	Log(string)
}

type Logger struct {
}

//为Logger添加Log()方法，同事实现Service的Log()方法
func (g *Logger) Log(l string) {

}

type GameService struct {
	Logger
}

//GameService的Start()方法实现了Service的Start()方法。
func (g *GameService) Start() {

}

//此时，实例化GameService，并将实例赋给Servic：
var s Service = new(GameService)
s.Start()
s.Log("hello")
//s就可以使用Start()方法和Log()方法，其中Startu由GameService实现，Log()方法由Logger实现

