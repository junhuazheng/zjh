//Go语言排序（借助sort.Interface接口）
// type interface interface {
// 	len() int

// 	Less(i, j int) bool

// 	Swap(i, j int)
// }

package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {

	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}

//由于将[]string定义为MyStringList类型，字符串切片初始化的过程等效于下面的写法：
// names := []string {
// 	"3. Triple Kill",
// 	"5. Penta Kill",
// 	"2. Double Kill",
// 	"4. Quadra Kill",
// 	"1. First Blood",
// }
