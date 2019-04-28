package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func testMap() {
	// var mmp map[string]interface{}
	mmp := make(map[string]interface{})

	mmp["username"] = "user"
	mmp["age"] = 19
	mmp["sex"] = "man"

	data, err := json.Marshal(mmp)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println("%s\n", string(data))
}

func main() {
	testMap()
	fmt.Println("----")
}
