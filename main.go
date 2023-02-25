package main

import (
	"GinChat/models"
	app "GinChat/router"
	"GinChat/utils"
	"fmt"
)

func main() {
	//初始化配置
	utils.ConfigFileInit()
	utils.MysqlInit()

	//测试获取数据库
	db := utils.GetDB()
	user := models.UserBasic{}
	db.Where("name = ?", "小徐").First(&user)
	fmt.Println("user: ", user)
	//调用app路由
	router := app.Router()
	router.Run(":8080") //指定端口号
}
