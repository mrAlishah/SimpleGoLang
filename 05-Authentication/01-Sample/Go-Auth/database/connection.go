package database

import (
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conection, err := gorm.Open(mysql.Open("root:alish@/Test"), &gorm.Config{})
	if err != nil {
		panic("could not connect to database!")
	}
	DB = conection
	conection.AutoMigrate(&models.User{})

}
