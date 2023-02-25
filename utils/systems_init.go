package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

// ConfigFileInit 读取配置文件
func ConfigFileInit() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config mysql:", viper.Get("mysql"))
}

// DB 定义全局的db对象，我们执行数据库操作主要通过他实现。
var DB *gorm.DB

func MysqlInit() {
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	fmt.Println("mysql.dns===", viper.GetString("mysql.dns"))
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
}

// GetDB 获取gorm db对象，其他包需要执行数据库查询的时候，utils.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return DB
}
