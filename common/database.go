package common

import (
	"awesomeProject/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//dirverName := viper.GetString("datasourse.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)
	db, err := gorm.Open(mysql.Open(args))
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
