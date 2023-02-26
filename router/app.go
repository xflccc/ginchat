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
	r.GET("/user/get-user-by-name", service.GetUserByName)
	r.GET("/user/get-all-user", service.GetAllUserList)
	r.GET("/user/create-user", service.CreateUser)
	r.DELETE("/user/delete-user", service.DeleteUser)
	r.PUT("/user/update-user", service.UpdateUser)

	return r
}
