//类型内嵌和结构体内嵌
//结构体允许其成员字段在声明时没有字段名而只有类型，这种形式的字段被称为类型内嵌或匿名字段类型内嵌：
type Data struct {
	//定义结构体中匿名字段
	int
	float32
	boll
}

ins := &Data{
	//将实例化的Data中的字段赋值
	int:     10,
	float32: 3.14,
	bool:    true,
}

//类型内嵌其实仍然拥有自己的字段名，只是字段名就是其类型本身而已
//结构体要求字段名必须唯一，因此一个结构体中同类型的匿名字段只能有一个

//结构体实例化后，如果匿名的字段类型为结构体，那么可以直接访问匿名结构体里的所有成员，这种方式被称为结构体内嵌
