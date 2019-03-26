//使用类型分支判断接口类型
//电子支付和现金支付：
package main

import "fmt"

type Alipay struct {
}

func (a *Alipay) CanUseFaceID() {

}

type Cash struct {
}

func (a *Cash) Stolen() {

}

type CantainUseFaceID interface {
	CanUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainUseFaceID:
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}

func main() {

	print(new(Alipay))

	print(new(Cash))
}
