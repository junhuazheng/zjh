// 多个值列表初始化结构体的格式：
ins := 结构体类型名{
	字段1的值,
	字段2的值,
	...
}

type Address struct {
	Province string
	City string
	Zipcode int
	phoneNumber string
}

addr := Address{
	"abc",
	"bcd",
	151515,
	"0",
}

fmt.Prinltn(addr)