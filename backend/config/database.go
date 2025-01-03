package config

import (
	"fmt"
	"log"

	"github.com/fauzan264/crowdfunding/backend/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var(
		DbHost = helper.GetEnv("DB_HOST", "db")
		DBPort = helper. GetEnv("DB_PORT", "3306")
		DbUser = helper.GetEnv("DB_USER", "root")
		DbPassword = helper.GetEnv("DB_PASSWORD", "root")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}