package main

import (
	"fmt"

	"github.com/kr/pretty"
)

type UserHasMany struct {
	//gorm.Model
	ID   uint32 `gorm:"primaryKey,autoIncrement"` //its name id
	Name string `gorm:"type:varchar(100);column:Name"`

	CreditCards []CreditCard `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreditCard struct {
	//gorm.Model
	ID     uint32 `gorm:"primaryKey,autoIncrement"` //its name id
	Number string `gorm:"type:varchar(10);column:Number"`
	UserID uint   `gorm:"column:UserID"`
}

//:: UserHasMany
func (user *UserHasMany) TableName() string {
	return "UserHasMany"
}

func (user *UserHasMany) ToString() string {
	return fmt.Sprint(pretty.Sprint(user))
}

func (user *UserHasMany) ToStringAll() string {
	return fmt.Sprint(pretty.Print(user))
}

//:: CreditCard
func (credit *CreditCard) TableName() string {
	return "CreditCard"
}

func (credit *CreditCard) ToString() string {
	return fmt.Sprintf("{ Id:%d \n  Number:%s \n } ", credit.ID, credit.Number)
}

//:: Other
func initSchemaUserHasMany() {
	// Migrate the schema
	//db.AutoMigrate(&User{})
	db.Migrator().DropTable(&UserHasMany{})
	db.Migrator().CreateTable(&UserHasMany{})
	db.Migrator().DropTable(&CreditCard{})
	db.Migrator().CreateTable(&CreditCard{})
	println("Table created")

}

func initDataUserHasMany() {
	//.Debug()
	db.Save(&UserHasMany{
		Name: "Mostafa",
		CreditCards: []CreditCard{
			{Number: "111"},
			{Number: "333"},
		},
	})

	db.Save(&UserHasMany{
		Name: "Hasan",
		CreditCards: []CreditCard{
			{Number: "222"},
		},
	})
}

func PreloadUserHasMany() {

	users := []UserHasMany{}

	// db.Preload("CreditCards").Find(&users)
	if err := db.Where("id = ? and \"Name\" = ?", 1, "Mostafa").Preload("CreditCards").Find(&users).Error; err != nil {
		fmt.Printf(err.Error())
	}
	printUsers(users)
}

func JoinsUserHasMany() {

	users := []UserHasMany{}
	tbl := `JOIN "CreditCard" ON "CreditCard"."UserID" = "UserHasMany"."id"`

	if err := db.Debug().Joins(tbl).Find(&users).Error; err != nil {
		fmt.Printf(err.Error())
	}
	printUsers(users)
}

func printUsers(users []UserHasMany) {
	for _, user := range users {
		fmt.Println(user.ToStringAll())
	}

}

func RowsUserHasMany() {

	rows, err := db.Debug().Table("UserHasMany").Where(`"UserHasMany"."Name" = ?`, "Mostafa").Joins(`JOIN "CreditCard" ON "CreditCard"."UserID" = "UserHasMany"."id"`).Select("*").Rows()
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer rows.Close()

	var user UserHasMany

	for rows.Next() {
		// ScanRows scan a row into user
		db.ScanRows(rows, &user)
		fmt.Println(user.ToStringAll())

	}
}

const (
	rawSQL = `
	SELECT 
	u."id",
	u."Name",
	c."id",
	c."Number",
	c."UserID" 
	
	FROM "UserHasMany" u
	LEFT JOIN "CreditCard" c ON c."UserID" = u.id
	`
)

func RawSqlUserHasMany() {

	rows, err := db.Debug().Raw(rawSQL).Rows()
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer rows.Close()

	users := make([]UserHasMany, 0)

	for rows.Next() {
		user := UserHasMany{}
		credit := CreditCard{}
		err = rows.Scan(&user.ID, &user.Name, &credit.ID, &credit.Number, &credit.UserID)
		if err != nil {
			fmt.Printf(err.Error())
		}

		user.CreditCards = append(user.CreditCards, credit)
		users = append(users, user)

		fmt.Println(user.ToStringAll())

	}
}
