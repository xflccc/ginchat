package service

import (
	"GinChat/models"
	"GinChat/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get index",
	})
}

func GetUserByName(c *gin.Context) {
	//测试获取数据库
	db := utils.GetDB()
	user := models.UserBasic{}
	db.Where("name = ?", "小徐").First(&user)

	c.JSON(200, gin.H{
		"name1":    user.Name,
		"password": user.PassWord,
		"ID":       user.ID,
	})
}

func GetAllUserList(c *gin.Context) {
	list := models.GetAllUserList()
	fmt.Println("list: ", list)
	c.JSON(200, gin.H{
		"userList": list,
	})
}
