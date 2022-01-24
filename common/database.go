package common

import (
	"Gowarden/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	fmt.Println("database init")
	InitDB()
}

func InitDB() *gorm.DB {
	host := config.MYSQL_HOST
	port := config.MYSQL_PORT
	database := config.MYSQL_DB
	username := config.MYSQL_USERNAME
	password := config.MYSQL_PASSWORD
	chaset := config.MYSQL_CHARSET

	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		username,
		password,
		host,
		port,
		database,
		chaset)

	fmt.Println("database url：", sqlStr)
	db, err := gorm.Open(mysql.Open(sqlStr))
	if err != nil {
		fmt.Println("Open database failed!", err)
		panic("Open database failed!" + err.Error())
	}
	DB = db
	return DB
}
