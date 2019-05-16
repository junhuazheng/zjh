package main

import (
	"fmt"
)

type Customer struct {
	Name    string
	Deposit float64
}

func TransferAccounts(name1, name2 Customer, money float64) (Customer, Customer, error) {

	if name1.Deposit-money < 0 {
		return name1, name2, fmt.Errorf("对不起, [%s]的用户余额已不足", name1.Name)
	} else {
		name1.Deposit = name1.Deposit - money
		name2.Deposit = name2.Deposit + money
	}

	return name1, name2, nil
}

func main() {
	zjh := Customer{"郑俊华", 10000}
	Mayun := Customer{"马芸", 1}

	a := 5000.00
	b := 20000.00

	name1, name2, err := TransferAccounts(zjh, Mayun, a)
	fmt.Printf("转账中，转账金额为:%v\n", a)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name1, name2)
		fmt.Println("转账操作成功！")
	}

	name1, name2, err = TransferAccounts(zjh, Mayun, b)
	fmt.Printf("转账中，转账金额为:%v\n", b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name1, name2)
		fmt.Println("转账操作成功！")
	}

}
