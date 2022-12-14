package db

/**
文档: https://gorm.io/zh_CN/docs/query.html
*/

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB = GetDB()
)

func GetDB() *gorm.DB {

	//host := util.GetCfgStr("database.host")
	//user := util.GetCfgStr("database.user")
	//password := util.GetCfgStr("database.password")
	//dbname := util.GetCfgStr("database.dbname")
	//port := util.GetCfgStr("database.port")

	host := "localhost"
	port := "5432"
	//user := "uname"
	//password := "zhenxun"
	//dbname := "testdb"

	user := "postgres"
	password := "bgy123456"
	dbname := "monitor"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	println("数据连接初始化完成")
	return db
}
