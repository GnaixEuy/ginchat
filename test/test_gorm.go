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
	user := &models.UserBasic{Name: "GnaixEuy"}
	db.Create(user)
	fmt.Println(db.First(user, 1))
}
