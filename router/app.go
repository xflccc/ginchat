package router

import (
	"GinChat/docs"
	"GinChat/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Router 请求路由/**
func Router() *gin.Engine {
	r := gin.Default()

	//swagger 配置
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//欢迎页
	r.GET("/index", service.GetIndex) //掉用service 层
	//user 表相关
	r.GET("get-user-by-name/index", service.GetUserByName)
	r.GET("get-all-user/index", service.GetAllUserList)

	return r
}
