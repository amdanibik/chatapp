package configs

import (
	"app/models/chats"
	"app/models/rooms"
	"app/models/users"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_DTBS string
}

func InitDB() {
	config := ConfigDB{
		DB_USER: "root",
		DB_PASS: "1234Abcd",
		DB_HOST: "127.0.0.1",
		DB_PORT: "3306",
		DB_DTBS: "chat",
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_DTBS)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&users.User{})
	DB.AutoMigrate(&rooms.Room{})
	DB.AutoMigrate(&chats.Chat{})
}
