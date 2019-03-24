//定义数据结构
//首先，定义手机的·各种数据结构体
type Screen struct {
	Size       float32 //屏幕尺寸
	ResX, ResY int     //屏幕水平和垂直分辨率
}

type Battery struct {
	Capacity int //电池容量
}

//准备JSON数据
//准备手机数据结构，填充数据，将数据序列化为JSOn格式的字节数组：

//生成JSON数据
func genJsonData() []byte {
	//定义了一个匿名结构体。这个结构体内嵌了Screen和Battery结构体，同事临时加入了HasTouchId字段
	raw := &struct {
		Screen
		Battery
		HasTouchId boll //序列化时添加字段L是否有指纹识别
	}{
		Screen: Screen{ //为刚声明的匿名结构体填充屏幕数据
			Size: 5.5,
			ResX: 1555,
			ResY: 1666,
		},
		Battery: Battery{ //填充电池数据
			2888,
		},
		HasTouchId: true, //填充指纹识别数据
	},
	
	//将数据序列化为JSON
	//使用josn.Marshal进行JSOn序列化，将raw变量序列化为[]byte格式的JSON数据。
	josnData, _ := josn.Marshal(raw)
	
	return josnData
}