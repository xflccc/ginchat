package service

import (
	"GinChat/models"
	"GinChat/utils"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tag 首页
// @Success 200 {string} hello
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "欢迎来到 index",
	})
}

// GetUserByName
// @Tag 根据用户名获取用户
// @Success 200 {string} hahaha
// @Router /get-user-by-name/index [get]
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

// GetAllUserList
// @Tag 获取所有用户
// @Success 200 {string} json {"code","userList"}
// @Router /get-all-user/index [get]
func GetAllUserList(c *gin.Context) {
	list := models.GetAllUserList()
	bytes, _ := json.Marshal(list) //集合转json
	fmt.Println("list:", string(bytes))
	c.JSON(200, gin.H{
		"userList": list,
	})
}
