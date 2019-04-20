package main

import (
	"io"
	"os"
)
/*
编程语言中反射的概念
在计算机科学领域，反射是指一类应用，他们能够自描述和自控制。
也就是说，这类应用通过采用某种机制对自己行为的描述（self-representation）和监测（examination）
并能根据自身行为的状态和结果，调整或修改应用所描述行为的状态和相关的语义。

每种语言的反射模型都不同，并且有鞋语言根本不支持反射。
Golang语言实现了反射，反射机制就是在运行时动态的调用对象的方法和属性，reflect包就是反射相关的，只要包含这个包就可以使用了。
Golang的gRPC也是通过反射实现的。

interface 和反射
Golang关于类型设计的一些原则：
1、变量包括（type，value）两部分
2、type包括static type和concrete type，简单来说static type就是在编码时看见的类型（如int、string）
concrete type 是runtime系统看见的类型。
3、类型断言能否成功，取决于变量的concrete type，而不是static type 
因此，一个reader变了那个如果它的concrete type也实现了write方法的话，它也可以被类型断言为writer。

反射就是建立在类型之上的，Golang的指定类型的变量的类型是静态的（也就是指定int、
string这些的变量，它的type是static type），在创建变量的时候就已经确定，反射主要与Golang的interface类型相关
（他的type是concrete type），只有interface类型才有反射一说。

在Golang的实现中，每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型：*/
(value， type)

/*
value是实际变量值，type是实际变量的类型。
一个interface{}类型的变量包含了两个指针，一个指针指向值得类型（对应concrete type），另外一个指针指向实际的值（对应value）

创建一个类型为*os.File变量，然后将其赋给一个接口变量r：*/
tty, err := os.OpenFile("/dev/tyy", os.O_RDWR, 0)

var r io.Reader

r = tty
/*接口变量r的pair中将记录如下信息：(tty, *os.File)，这个pair在接口变量的连续赋值过程中是不变的

将接口变量r赋给另一个接口变量w：*/
var w io.Writer
w = r.(io.Writer)