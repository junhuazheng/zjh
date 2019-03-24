package main

//分离JSOn数据
//调用genJsonData获得JSon数据，将需要的字段填充到匿名结构体实例中
//通过json.Unmarshal反序列化Json数据打成分离Json数据效果：

func main() {

	//调用genJsonData()函数，获得[]byte类型的JSON数据
	josnData := genJsonData()

	fmt.Println(string(jsonData))

	//只需要屏幕和指纹识别的结构体和实例
	screenAndTouch := struct {
		Screen
		HasTouchId bool
	}{} //{}表示将结构体实例化

	//调用json.Unmarshal，输入完整的JSON数据（jsonData)，将数据按前面定义的结构体格式序列化到screenAndTouch中。
	json.Unmarshal(jsonData, &screenAndTouch)

	//输出screenAndTouch的详细结构
	fmt.Printf("&+v\n", screenAndTouch)

	//只需要电池和指纹识别信息的结构体和实例
	batteryAndTouch := struct {
		Battery
		HasTouchId bool
	}{}

	//反序列化
	json.Unmarshal(jsonData, &batteryAndTouch)

	fmt.Prinltf("%+v\n", batteryAndTouch)
}
