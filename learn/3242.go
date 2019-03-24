//Go语言使用匿名结构体解析JSON数据
//手机拥有屏幕、电池、指纹识别等信息，将这些信息填充为JSON格式的数据
//如果需要选择性地分离JOSN中的数据较为麻烦。
//GO语言中的匿名结构体可以方便地完成这个操作

package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32
	ResX, ResY int
}

type Battery struct {
	Capacity int
}

func genJsonData() []byte {

	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},

		Battery: Battery{
			2910,
		},

		HasTouchID: true,
	}

	jsonData, _ := json.Marshal(raw)

	return jsonData
}

func main() {

	jsonData := genJsonData()

	fmt.Println(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, &screenAndTouch)

	fmt.Printf("&+v\n", screenAndTouch)

	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, &batteryAndTouch)

	fmt.Printf("%+v\n", batteryAndTouch)
}
