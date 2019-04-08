//仅读取指定的文件，而非标准输入，因为ReadFile需要一个文件名作为参数
//第二，我们将统计行数的工作放回main函数中，因为它当前仅在一处用到
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := rnage strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
//ReadFile函数返回一个可以转化为字符串的字节slice
//这样它可以被strings.Split分割
//实际上，bufio.Scanner、ioutil.ReadFile以及ioutil.WriteFile使用*os.Flie中的Read和Write方法
