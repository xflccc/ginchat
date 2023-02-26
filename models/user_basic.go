package models

import (
	"GinChat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	PassWord      string    `json:"passWord"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Identity      string    `json:"identity"`
	ClientIp      string    `json:"clientIp"`
	LoginTime     time.Time `json:"loginTime"`
	HeartbeatTime time.Time `json:"heartbeatTime"`
	LoginOutTime  time.Time `json:"loginOutTime"`
	IsLoginOut    bool      `json:"isLoginOut"`
	DeviceInfo    string    `json:"deviceInfo"`
}

// TableName 返回数据库名称的方法 gorm生成数据表的名称
func (table *UserBasic) TableName() string {
	return "user_basic"
}

// GetAllUserList (函数) 获取所有用户集合
func GetAllUserList() []*UserBasic {
	userList := make([]*UserBasic, 10)
	db := utils.GetDB()
	db.Find(&userList)
	return userList
}

// CreateUser (函数) 创建用户
func CreateUser(user *UserBasic) *gorm.DB {
	db := utils.GetDB()
	return db.Create(&user)
}

// DeleteUser (函数) 删除用户
func DeleteUser(user *UserBasic) *gorm.DB {
	db := utils.GetDB()
	return db.Delete(&user)
}

// UpdateUser (函数) 修改用户
func UpdateUser(user *UserBasic) *gorm.DB {
	db := utils.GetDB()
	return db.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone})
}
