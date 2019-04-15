/*有事吼，我们希望通过除了自然排序以外的其他方式对集合进行排序
例如，假设我们想按字符串的长度而不是字母的顺序对字符串进行排序。

为了使用Go语言红的自定义函数进行排序，我们需要一个相应的类型。
这里创建了一个ByLength类型，它只是内置[]string类型的别名。

需要实现 sort.Interface - len ， Less 和 Swap - 在这个类型上
所以可以使用sort包中的一般Sort函数。
Len和Swap通常在类型之剑是相似的，Less保存实际的自定义排序逻辑。
在这个例子中，要按照字符串长度的增加顺序排序，因此在河里使用lem(s [i]) 和 len(s [j])。*/

package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

/*所有这些都到尾吼，现在可以通过原石fruits切片转换为ByLength来实现在定义排序
然后对该类型切片使用sort.Sort()方法。

运行程序根据需要显示按照字符串长度排序的列表。

通过遵守创建自定义类型的魔石，在该类型上实现三个Interface方法
然后在自定义类型的集合上调用sort.Sort，可以通过任意函数对Go切片进行排序。*/
