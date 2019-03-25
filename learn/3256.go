//1、字符串切片的便捷排序
//sort包中有一个StringSlice类型，定义如下：
package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//Sort is a convenience method.
func (p StringSlice) Sort() {
	Sort(p)
}

//sort包中的StringSliece的代码与MyStringList的实现代码几乎一样。
//因此只需要使用sort包的StringSlice就可以更简单快速地进行字符串排序
names := sort.StringSlice{
	"3. Triple Kill",
	"5. Penta Kill",
	"2. Double Kill",
	"4. Quadra Kill",
	"5. First Blood",
}

sort.Sort(names)
