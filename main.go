package main

import (
	app "GinChat/router"
	"GinChat/utils"
)

func main() {
	//初始化配置
	utils.ConfigFileInit()
	utils.MysqlInit() // gorm 框架实现

	//调用app路由(gin框架实现)
	router := app.Router()
	router.Run(":8080") //指定端口号
}
