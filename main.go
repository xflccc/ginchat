package main

import app "GinChat/router"

func main() {
	//调用app路由
	router := app.Router()
	router.Run(":8080") //指定端口号
}
