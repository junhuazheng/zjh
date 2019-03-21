//接收器的格式如下：
func (接收器变量 接收器类型) 方法名(参数列表) 返回参数 {
	函数体
}

//例如：
func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}