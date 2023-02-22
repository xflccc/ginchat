package main

import (
	"GinChat/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	// 绑定实体类模板(UserBasic) 数据库没有表时自动创建
	db.AutoMigrate(models.UserBasic{})

	//获取user对象
	user := &models.UserBasic{}
	user.Name = "小徐"
	user.PassWord = "123"
	//user := &models.UserBasic{}
	//user.Name = "小徐"
	//user.PassWord = "123"

	//creat
	db.Create(user)

	//query
	query1 := db.First(user, 1) //根据主键查找
	fmt.Println(query1)
	query2 := db.First(user, "name = ?", "小徐") //根据姓名查找
	fmt.Println(query2)
}
