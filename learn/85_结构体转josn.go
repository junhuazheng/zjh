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

func testStruct() {
	uesr1 := &User{
		UserName: "user1",
		NickName: "上课看书",
		Age:      18,
		Birthday: "2028/9/9",
		Sex:      "男",
		Email:    "mayun@qq.com",
		Phone:    "911",
	}

	data, err := json.Marshal(uesr1)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func main() {
	testStruct()
	fmt.Println("----")
}
