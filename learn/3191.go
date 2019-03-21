//对数据的操作进行多步骤的处理被称为链式处理
package main

import (
	"fmt"
	"strings"
)

func StringProccess(list []string, chain []func(string) string) {

	for index, str := range list {

		result := str

		for _, proc := range chain {

			result = proc(result)
		}

		list[index] = result
	}
}

func removePrefix(str string) string {

	return strings.TrimPrefix(str, "go")
}

func main() {

	list := []string{
		"go scanner",
		"go parser",
		"go caompiler",
		"go printer",
		"go formater",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProccess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}
}
