package models

import (
	"GinChat/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLoginOut    bool
	DeviceInfo    string
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
	for _, value := range userList {
		fmt.Println(value)
	}
	return userList
}
