/*nil在Go语言中只能赋值给指针和借口。接口在底层的视线有两个部分：type 和 data。
在源码红，显式地将nil赋值给接口时，接口的type和data都将为nil。
此时，接口与nil值判断是相等的。
但如果将一个带有类型的nil复制给接口时，只有data为nil，而type为nil，此时接口与nil判断将不相等。

使用Mylmplement()函数实现fmt包中的Stringer接口
在GetStringer()函数中将返回这个接口。通过*Mylmplement指针变量置为nil提供GetStringer的返回值。
在main()中，判断GetStringer与nil是否相等：*/
package main

import "fmt"

type Mylmplement struct{}

func (m *Mylmplement) String() string {
	return "abc"
}

func GetStringer() fmt.Stringer {

	var s *Mylmplement = nil

	return s
}

func main() {

	//判断返回值是否为nil
	if GetStringer() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GerStringer() != nil")
	}
}

//虽然接口里的value为nil，但type带有*Mylmplement信息，使用==判断相等时，依然不为nil。
