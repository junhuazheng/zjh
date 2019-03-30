package main

import (
	"github.com/astaxie/beego"
)

var (
	mapString map[string]string
)

type AddController struct {
	beego.Controller
}

func (this *AddController) Get() { //get post  == http

	mapString["tmp"] = "俊华"
	this.Ctx.WriteString("hello world  ok")

}

type MainController2 struct {
	beego.Controller
}

func (this *MainController2) Get() {
	str := mapString["tmp"]
	this.Ctx.WriteString(str)
}

func main() {
	mapString = make(map[string]string)
	beego.Router("/add", &AddController{})
	beego.Router("/get", &MainController2{})
	beego.Run()
}
