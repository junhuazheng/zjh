package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Fangfa() {
	fmt.Println("Allen.Wu Fangfa")
}

func main() {
	user := User{1, "Allen.Wu", 25}

	DofiledAndMethod(user)
}

func DofiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)              //获取Field
		value := getValue.Field(i).Interface() //得到对应的value
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)                 //获取interface的方法
		fmt.Printf("%s: %v\n", m.Name, m.Type) //reflect.Method的m.Name为方法名，m.Type为方法类型
	}
}
