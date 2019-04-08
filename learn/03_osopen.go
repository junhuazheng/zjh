//许多程序既可以像dup一样从标准输入进行读取，也可以从具体的文件读取。
//下一个版本的dup程序可以从标准输入或者一个文件列表进行读取，使用os.Open函数来逐个打开：
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) = 0 {
		countLines(os.Stdin, counts)
	}else {
		for _, arg  := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
