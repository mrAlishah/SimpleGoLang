package main

import "fmt"

type User struct {
	//gorm.Model
	ID        uint64 `json:"id"  gorm:"primaryKey,autoIncrement"`
	UserName  string `json:"username" gorm:"type:varchar(100)"`
	FirstName string `json:"firstname" gorm:"type:varchar(100)"`
	LastName  string `json:"lastname" gorm:"type:varchar(100)"`
	Salary    uint   `json:"Salary" gorm:"default:0"`
}

var users []User = []User{
	User{UserName: "111", FirstName: "ali", LastName: "feyzi"},
	User{UserName: "222", FirstName: "mahtab", LastName: "ebrahimi"},
	User{UserName: "333", FirstName: "ali", LastName: "aslani"},
	User{UserName: "444", FirstName: "mostafa", LastName: "alishah"},
}

func initUser() {
	// Migrate the schema
	//db.AutoMigrate(&User{})
	db.Migrator().DropTable(&User{})
	db.Migrator().CreateTable(&User{})
	println("Table created")

}

func introCRUD() {
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
	printAll()

	// Delete - delete user
	db.Delete(&User{}, 3)
	printAll()

}

func printAll() {
	db.Find(&users)
	for _, user := range users {
		fmt.Printf("Record  %v \n", user)
	}
	println("\n==============================\n")
}

// NotFound checks if a record exists in the database
func (u *User) NotFound() bool {
	return u.ID == 0
}
