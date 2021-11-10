package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresCnf struct {
	HOST     string
	PORT     int16
	USER     string
	PASSWORD string
	DBNAME   string
	SSLMODE  string
}

type User struct {
	//gorm.Model
	ID        uint64 `json:"id"  gorm:"primaryKey,autoIncrement"`
	UserName  string `json:"username" gorm:"type:varchar(100)"`
	FirstName string `json:"firstname" gorm:"type:varchar(100)"`
	LastName  string `json:"lastname" gorm:"type:varchar(100)"`
}

var users []User = []User{
	User{UserName: "111", FirstName: "ali", LastName: "feyzi"},
	User{UserName: "222", FirstName: "mahtab", LastName: "ebrahimi"},
	User{UserName: "333", FirstName: "ali", LastName: "aslani"},
	User{UserName: "444", FirstName: "mostafa", LastName: "alishah"},
}

func main() {
	srvConfig := postgresCnf{
		HOST:     "localhost",
		PORT:     5432,
		USER:     "postgres",
		PASSWORD: "alish",
		DBNAME:   "postgres",
		SSLMODE:  "disable",
	}
	dsn := connectionStringPostgres(srvConfig)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	println("Connection to database established")

	// Migrate the schema
	//db.AutoMigrate(&User{})
	db.Migrator().DropTable(&User{})
	db.Migrator().CreateTable(&User{})
	println("Table created")

	for _, user := range users {
		db.Create(&user)
		fmt.Printf("inserted %v \n", user)
	}
	println("inserted ")
	println("\n==============================\n")
	// Read
	user := User{}
	db.First(&user, 1) // find user with integer primary key
	fmt.Printf("First %v \n", user)

	user = User{} //must be null
	//db.Where("user_name = ?", "222").First(&user)
	db.First(&user, "user_name = ?", "222") // find user with username 222
	fmt.Printf("First user_name = ? %v \n", user)
	println("\n==============================\n")

	// Update - update user's lastname to akbari
	db.Model(&user).Update("user_name", "asgari")

	// Update - update multiple fields
	db.Model(&user).Updates(User{UserName: "555", FirstName: "masoud"}) // non-zero fields
	//db.Model(&user).Updates(map[string]interface{}{"UserName": "555", "FirstName": "masoud"})
	printAll(db)

	// Delete - delete user
	db.Delete(&User{}, 3)
	printAll(db)
}

func connectionStringPostgres(config postgresCnf) string {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s",
		config.HOST, config.PORT, config.USER, config.PASSWORD, config.DBNAME, config.SSLMODE)
}

func printAll(db *gorm.DB) {
	db.Find(&users)
	for _, user := range users {
		fmt.Printf("Record  %v \n", user)
	}
	println("\n==============================\n")
}
