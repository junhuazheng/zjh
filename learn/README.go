什么是interface
   简单的说，interface是一组method的组合，我们通过interface来定义对象的一组行为。
换句话说，一个interface类型定义了一个“方法集合”作为其接口。interface类型的变量可以
保存含有属于这个interface类型的任何类型的值，这时我们就说这个类型实现了这个接口，未被
初始化的interface类型变量的零值为空。

interface类型和值
   接口类型实际上是一组method签名的清单。interface类型定义了一组方法，如果某个对象实
现了某个接口的所有方法，则此对象就实现了此接口。接口也是一种数据类型。如果你声明了一个接口
变量，这变量能够存储任何实现该接口的对象类型。最后，任意的类型都实现了空interface（我们
这样定义：interface{}），也就是包含0个method的interface。