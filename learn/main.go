package main

import (
	"time"
	"fmt"
	"DEMOORM/orm"
	_"github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	TableName orm.TableName "userinfo"
	UserName string `name:"username"`
	Uid int `name:"uid"PK:"true"auto:"true"`
	DepartName string`name:"departname"`
	Created string `name:"created"`
}

func  main()  {
	ui := UserInfo{UserName:"CHAIN", DepartName:"TEST", Created:time.Now().String()}
	orm.Register(new(UserInfo))
	db, err := orm.NewDb("mysql", "root:pwd@tcp(locallhost:3306/mysql?charset=utf8")
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return
	}
	defer db.Close()

	//插入测试
	err = db.Insert(&ui)
	if err != nil {
		fmt.Println("插入时错误: ", err.Error())
	}
	fmt.Println("插入成功")

	//修改测试
	ui.UserName = "AAAA"
	err = db.Update(ui)
	if err != nil {
		fmt.Println("修改时错误: ", err.Error())
	}
	fmt.Println("修改成功")

	//删除测试
	err = db.Delete(ui)
	if err != nil {
		fmt.Println("删除时错误: ", err.Error())
	}
	fmt.Println("删除成功")

	//查询测试
	res, err := db.From("userinfo").
	Select("username", "departname", "uid").
	Where("uid_gt", 20).
	Where("username", "chain").Get()
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println(res)
}