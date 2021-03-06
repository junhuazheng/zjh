//接口被实现的条件二：接口中所有方法均被实现
//当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用
//为DataWriter添加一个方法：
type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}

//新增CanWrite()方法，返回bool。此时在编译代码，报错：
//需要在file中实现CanWrite()方法才能正常使用DataWriter()。

//Go语言的借口实现是隐式的，无须让实现接口的类型写出实现了哪些接口。这个设计被称为非入侵式设计

//实现者在编写方法时，无法预测未来哪些方法会变为接口。
//一旦某个接口创建出来，要求旧的代码来实现这个接口时，就需要修改旧的代码的派生部分，这一版会造成雪崩式的重新编译

//传统的派生式接口及类关系构建的模式，让类型间拥有强耦合的斧子关系。
//对于Go语言来说，非入侵式设计让实现者的所有类型均是平行的、组合的。
//如何组合则留到使用者编译时再确认。
//因此，不需要同事也不可能有“类派生图“，开发者唯一需要关注的就是”我需要什么？“，以及“我能实现什么？”