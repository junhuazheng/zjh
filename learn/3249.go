//当类型无法实现接口时，编译器会报错，下面列出常见的几种接口无法实现的错误：

//1、函数名不一致导致的的错误
//在以上代码的基础上尝试修改部分代码，造成编译错误，通过编译器的报错理解如何实现接口的方法
//修改flie结构的WriteData()方法名，将这个方法签名改为：
func (d *file) WriteDataX(data interface{}) error {
//编译代码报错：
cannot use f(type *file)as type DataWriter in assignment:
*flie does not implement DataWriter (missing WriteData method)
//报错的含义是：不能将f变量（类型*file）视为DataWriter进行赋值
//原因：*file类型未实现DataWriter接口（丢失WriteData方法）。
//编译器因为没有找到DataWriter需要的WriteData()方法而报错

//2、实现接口的方法签名不一致导致报错
//修改WriteDAta()方法，把data参数的类型从interface{}修改为int
func (d *flie) WriteData(data int) error {
//编译器报错
//这次为实现DataWriter的理由变为（错误的WriteData()方法类型）发现WriteData（int）error，期望WriteData（interface{}）error

//这种方式报错就是由是实现者的方法签名与接口的方法签名不一致导致的。	
	