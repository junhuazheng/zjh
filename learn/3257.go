//2、对整型切片进行排序
//除了字符串可以使用sort包进行便捷排序外，还可以使用sort.IntSlice进行整型切片的排序
//sort.IntSlice的定义如下：
package main

import (
	"fmt"
	"sort"
)

typ IntSlice []int

func (p IntSlice) len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p IntSlice) Sort() {
	Sort(p)
}