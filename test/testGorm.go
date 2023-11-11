package main

import (
	"fmt"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("failed to connect database")
	}
	//迁移 schema
	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		return
	}

	// Create
	user := &models.UserBasic{
		Name:          "GanixEuy",
		Email:         "test@gnaixeuy.cn",
		LogOutTime:    0,
		IsLogOut:      false,
		Password:      "test",
		ClientPort:    "test",
		LoginTime:     0,
		ClientId:      "test",
		DeviceInfo:    "test",
		HeartbeatTime: 0,
		Phone:         "test",
	}
	db.Create(user)
	fmt.Println(db.First(user, 1))
}
