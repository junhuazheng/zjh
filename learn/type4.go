// 键值对初始化结构体格式：
ins := 结构体类型名{
	字段1:字段1的值,
	字段2:字段2的值,
	...
}

type People struct {
	name string
	child *People
}

relation := &People{
	name:"爷爷",
	child:&People{
		name:"爸爸",
		child:&People{
			name:"我",
		},
	},
}