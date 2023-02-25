package router

import (
	"GinChat/service"

	"github.com/gin-gonic/gin"
)

// Router 请求路由/**
func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex) //掉用service 层

	//user 表相关
	r.GET("get-user-by-name/index", service.GetUserByName)
	r.GET("get-all-user/index", service.GetAllUserList)

	return r
}
