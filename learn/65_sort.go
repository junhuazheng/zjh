type interface
type Interface interface {
	Len() int //Len为集合内元素的总数
	Less(i, j int) bool //如果index为i元素小于index为j的元素，则返回ture，否则返回false
	Swap(i, j int) //Swap交换索引为i和j的元素
}

//sort相关的一些方法：

func Float64s(a []float64) //将类型为float64的silce a 以升序方式进行排序

func Float64sAreSorted(a []float64) bool //判断是否已经进行排序

func Ints(a []int) //以升序排列int切片

func IntsAreSorted(a []int) bool //判断int切片是否已经按升序排列。

func IsSorted(data interface) bool //判断数据是否已经排序，包括各种可sort的数据类型的判断

func Strings(a []string) //以升序排列string切片

func StringsAreSortde(a []string) bool //判断string切片是否已经按升序排列。

func Sort(data interface)
//对data进行排序（不保证相等元素的相对顺序不变）data默认为升序，执行Reverse后为降序

func Reverse(data interface) interface
//将data的排序动作更改为降序，Reverse并不改变元素顺序，只改变排序行为。
//更改操作不可逆，更改后的对象不可以再次Reverse。