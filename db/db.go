package db

import (
	"log"
	"zhihu-column-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库操作对象
var DB *gorm.DB

// ConnectDB 连接数据库
func ConnectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:hu123456@tcp(127.0.0.1:3306)/zhihu-column?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&model.User{})

}
