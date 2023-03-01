package service

import (
	"GinChat/models"
	"GinChat/utils"
	"encoding/json"
	"fmt"

	"strconv"

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
// @Tag 用户基本信息
// @Summary 根据用户名获取用户
// @Param name query string false "用户名称" maxlength(100)
// @Success 200 {string} hahaha
// @Router /user/get-user-by-name [get]
func GetUserByName(c *gin.Context) {
	//测试获取数据库
	db := utils.GetDB()
	user := models.UserBasic{}
	db.Where("name = ?", c.Query("name")).First(&user)

	c.JSON(200, gin.H{
		"name1":    user.Name,
		"password": user.PassWord,
		"ID":       user.ID,
	})
}

// GetAllUserList
// @Tag 用户基本信息
// @Summary 获取所有用户
// @Success 200 {string} json {"code","userList"}
// @Router /user/get-all-user [get]
func GetAllUserList(c *gin.Context) {
	list := models.GetAllUserList()
	bytes, _ := json.Marshal(list) //集合转json
	fmt.Println("list:", string(bytes))
	c.JSON(200, gin.H{
		"userList": list,
	})
}

// CreateUser
// @Tag 用户基本信息
// @Summary 注册用户
// @Param name query string false "用户名称" maxlength(100)
// @Param password query string false "密码" maxlength(100)
// @Param passwordv2 query string false "密码" maxlength(100)
// @Param phone query string false "手机号" maxlength(100)
// @Success 200 {string} json {"code","userList"}
// @Router /user/create-user [get]
func CreateUser(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	passwordv2 := c.Query("passwordv2")
	phone := c.Query("phone")
	if password != passwordv2 {
		c.JSON(500, gin.H{
			"message": "密码不一样！",
		})
		return
	}
	//校验手机哈是否合法
	if !utils.VerifyMobileFormat(phone) {
		c.JSON(500, gin.H{
			"message": "手机号格式错误，请重新输入！",
		})
		return
	}
	//查看用户名是否已经存在
	oldUser := models.GetUserByName(name)
	if oldUser.ID != 0 {
		c.JSON(500, gin.H{
			"message": "用户名已经存在，请重新输入！",
		})
		return
	}
	//判断手机号码是否存在
	if models.GetUserByPhone(phone).ID != 0 {
		c.JSON(500, gin.H{
			"message": "手机号码已经存在，请重新输入！",
		})
		return
	}
	user := models.UserBasic{}
	user.Name = name
	user.PassWord = password
	user.Phone = phone
	models.CreateUser(&user) //dao
	c.JSON(200, gin.H{
		"message": "添加用户成功",
	})
}

// DeleteUser
// @Tag 用户基本信息
// @Summary 删除用户
// @Param id query string true "用户id" maxlength(100)
// @Success 200 {string} json {"code","userList"}
// @Router /user/delete-user [delete]
func DeleteUser(c *gin.Context) {
	query := c.Query("id")
	id, _ := strconv.Atoi(query)
	user := models.UserBasic{}
	user.ID = uint(id)
	models.DeleteUser(&user) //dao
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Tag 用户基本信息
// @Summary 修改用户
// @Param id formData  string true "用户id" maxlength(100)
// @Param name formData  string true "用户名称" maxlength(100)
// @Param password formData  string true "用户密码" maxlength(100)
// @Param phone formData  string true "用户手机号" maxlength(100)
// @Success 200 {string} json {"code","userList"}
// @Router /user/update-user [put]
func UpdateUser(c *gin.Context) {
	query := c.PostForm("id")
	id, _ := strconv.Atoi(query)
	user := models.UserBasic{}
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	models.UpdateUser(&user) //dao
	c.JSON(200, gin.H{
		"message": "修改用户成功",
	})
}
