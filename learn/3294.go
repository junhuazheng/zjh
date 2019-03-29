/*Go语言单向通道————通道中的单行道
Go的通道可以在声明时约束其操作方向，如只发送或是只接受。
这种被约束方向的通行到被称为单向通道

1、单向通道的声明
只能发送的通道类型为cahn<-，只能接受的通道类型为<-chan,格式如下：*/
var 通道实例 chan<- 元素类型 //只能发送通道
var 通道实例 <-chan 元素类型 //只能接受通道
/*元素类型：通道包含的元素类型
  通道实例： 声明的通道变量

  2、单向通道的使用例子*/
ch := make(chan int)

//声明一个只能发送的通道类型，并赋值为ch
var chSendOnly chan<- int = ch

//声明一个只能接收的通道类型，并赋值为ch
var chRecvOnly <-chan int = ch

//使用make创建通道时，也可以创建一个只发送或只读取的通道：
ch := make(<-chan int)

var chReadOnly <-chan int = ch
<-chRedOnly
//但是，一个不能填充数据（发送）只能读取的通道是毫无意义的

/*3、time包中的单向通道
  time包中的计时器会返回一个timer实例，代码如下：*/
timer := time.NewTimer(time.Second)

//time的Timer类型定义如下：
type Timer struct {
	C <-chan Time
	r runtimeTimer
}
/* C <-chan Time中的C通道的类型就是一种只能接受的单向通道。
   如果此处不进行通道方向约束，一旦外部向通道发送数据，将会造成其他使用到计时器的地方逻辑产生混乱
   因此，单向通道有利于代码接口的严谨性*/                                                                                                                                                                                                                  