//JSON格式数据，用匿名结构体选择性地分离JSON中的数据。
package main

import (
	"encoding/json"
	"fmt"
)

type Scr struct {
	Size     float32
	ReX, ReY int
}

type Bat struct {
	Capa int
}

//生成json数据
func genJsonData() []byte {

	raw := &struct {
		Scr
		Bat
		Has bool
	}{
		Scr: Scr{
			Size: 5.5,
			ReX:  1920,
			ReX:  1080,
		},

		Bat: Bat{
			2910,
		},
		HasT: true,
	}

	jsonData, _ := json.Marshal(raw)

	return jsonData
}

func main() {
	jsonData := genJsonData()

	fmt.Println(string(jsonData))

	scrA := struct {
		Scr
		Has bool
	}{}

	json.Unmarshal(jsonData, &scrA)

	fmt.Printf("%+v\n", scrA)

	batA := struct {
		Bat
		Has boll
	}{}

	json.Unmarshal(jsonData, &batA)

	fmt.Printf("%+v\n", batA)
}
