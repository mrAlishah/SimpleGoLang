package main

import (
	"fmt"

	"gorm.io/gorm"
)

//belong to =  `User` belongs to `Calender`, `CalenderID` is the foreign key
//Has One =   User has one Calender, CalenderID is the foreign key
type User1 struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(100);column:UserName"`
	FirstName string `gorm:"type:varchar(100);column:FirstName"`
	LastName  string `gorm:"type:varchar(100);column:LastName"`

	//has one ~ uniqe every calenderId=UserID  and ownership is user// use belong to too
	Calender Calender `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //OnDelete:SET NULL;

	//Calender   Calender `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //belong to
	//CalenderID uint           `gorm:"column:CalenderID"`                                           //belong to ~ more multi user=CalenderID and ownership is calender , create record Calender first
}

type Calender struct {
	gorm.Model
	Name   string `gorm:"column:Name"`
	UserID uint   `gorm:"column:UserID"` //has one
}

//:: UserHasOne
func (user *User1) TableName() string {
	return "User1"
}

func (user *User1) ToString() string {
	return fmt.Sprintf("{ Id:%d \n  UserName:%s \n  FirstName:%s \n  LastName:%s \n }", user.ID, user.UserName, user.FirstName, user.LastName)
}

func (user *User1) ToStringAll() string {
	return fmt.Sprintf("{ Id:%d \n  UserName:%s \n  FirstName:%s \n  LastName:%s \n \t\t Calender:{Id:%d \n\t\t Name:%s \n\t\t UserId:%d \n\t\t} \n}",
		//user.ID, user.UserName, user.FirstName, user.LastName, user.Calender.ID, user.Calender.Name, user.Calender.UserID) //has one
		user.ID, user.UserName, user.FirstName, user.LastName, user.Calender.ID, user.Calender.Name, user.Calender.ID) //has one
	//user.ID, user.UserName, user.FirstName, user.LastName, user.Calender.ID, user.Calender.Name, user.CalenderID) //belong to
}

//:: CalenderHasOne
func (cal *Calender) TableName() string {
	return "Calender"
}

func (cal *Calender) ToString() string {
	return fmt.Sprintf("{ Id:%d \n  Name:%s \n } ", cal.ID, cal.Name)
}

//:: Other
func initUserHasOne() {
	// Migrate the schema
	//db.AutoMigrate(&User1{})
	db.Migrator().DropTable(&User1{})
	db.Migrator().CreateTable(&User1{})
	db.Migrator().DropTable(&Calender{})
	db.Migrator().CreateTable(&Calender{})
	println("Table created")

}

func initHasOne() {
	//.Debug()
	db.Save(&User1{
		UserName: "Mostafa",
		Calender: Calender{
			Name: "Import Events",
		},
	})

	db.Save(&User1{
		UserName: "Hasan",
		Calender: Calender{
			Name: "Import Events1",
		},
	})

	u := &User1{}
	//db.Preload(clause.Associations).First(&u)
	//db.Preload("Calender").First(&u)
	//fmt.Println(u.ToStringAll())
	db.Joins("Calender").First(&u)
	fmt.Println(u.ToStringAll())
}
