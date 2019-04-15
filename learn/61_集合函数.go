/*我们经常需要在程序中对数据集合执行操作
例如选择满足给定谓词的所有项目，或将所有项目映射到具有自定义函数的新集合。
（谓词：用来描述或排定客体性质、特征或者客体之间关系的词项）

在某些语言中，通用数据结构和算法是管用的。Go不支持泛型；
在Go中，如果并且当他们对于程序和数据类型特别需要时，提供集合函数是常见的。

这里是一些字符串切片的示例收集函数。
可以使用这些示例来构建自己的函数。
注意，在某些情况下，直接内联集合操作代码可能是最清楚的，而不用创建和调用辅助函数。*/

package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	f i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string)bool)bool {
	f _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(styring) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string)bool) []string {
	vsf := make([]string, 0)
	for _, v :=range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}

/*
返回目标字符串t的第一个索引，如果未找到匹配，则返回-1。如果目标字符串t在切片中，则返回true。
如果切片中的一个字符串满足谓词f，则返回true。
如果切片中的所有字符串都满足谓词f，则返回true。
返回包含切片中满足谓词f的所有字符串的新切片。
返回一个新的切片，包含将函数f应用于原始切片中的每个字符串的结果
*/


