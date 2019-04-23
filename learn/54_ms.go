package main

import (
	"errors"
	"fmt"
)

var ErrDidNotWrork = errors.New("did not wrok")

// func DoTheThing(reallyDoIt bool) (err error) {
// 	if reallyDoIt {
// 		result, err := tryTheThing()
// 		if err != nil || result != "it wroked" {
// 			err = ErrDidNotWrork
// 		}
// 	}
// 	return err
// }

func tryTheThing() (string, error) {
	return "", ErrDidNotWrork
}

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}

/*
变量作用域
输出结果为:
<nil>
<nil>
因为if语句块内的err变量会遮罩函数作用域内的err变量
*/
//修改为：
func DoTheThing(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it wroked" {
			err = ErrDidNotWrork
		}
	}
	return err
}
