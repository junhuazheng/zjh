package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Deposit float64
}

//自定义的错误
func CustomError() error {
	return errors.New("对不起，您的余额已不足。")
}

//用来实现转账功能的函数
func TransferAccounts(name1, name2 Customer, money float64) (Customer, Customer, error) {
	if name1.Deposit-money < 0 {
		return name1, name2, CustomError()
	} else {
		name1.Deposit = name1.Deposit - money
		name2.Deposit = name2.Deposit + money
	}
	return name1, name2, nil
}

func main() {

	a := 50000.00
	b := 800000.00

	zjh := Customer{"郑俊华", 100000}
	Mayun := Customer{"玛云", 100}

	name1, name2, err := TransferAccounts(zjh, Mayun, a)
	fmt.Printf("正在转账中，转账金额为:%v\n", a)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name1, name2) //显示余额
		fmt.Println("转账操作成功！")
	}

	name1, name2, err = TransferAccounts(zjh, Mayun, b)
	fmt.Printf("正在转账中，转账金额为:%v\n", b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name1, name2)
		fmt.Println("转账操作成功！")
	}
}
