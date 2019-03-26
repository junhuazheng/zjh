//Go语言使用空接口实现可以保存任意值的字典

//空接口可以保存任何类型这个特性可以方便地用于容器的设计

package main

import "fmt"

type Dictionary struct {
	data map[interface{}]interface{}
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

func (d *Dictionary) Set(key interface{}, vlue interface{}) {
	d.data[key] = value
}

func (d *Dictionary) visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

func NewDictionart() *Dictionary {
	d := &Dictionary{}

	d.Clear()
	return d
}

func main() {
	dict := NewDictionart()

	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)

	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)

	dict.visit(func(key, value interface{}) bool {

		if value.(int) > 40 {
			fmt.Println(key, "is expensive")
			return true
		}
		fmt.Println(key, "is cheap")

		return true
	})
}
