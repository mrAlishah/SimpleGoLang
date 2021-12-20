package main

import (
	"fmt"

	"github.com/kr/pretty"
)

type UserM struct {
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
func (user *UserM) TableName() string {
	return "UserM"
}

func (user *UserM) ToString() string {
	return fmt.Sprint(pretty.Sprint(user))
}

func (user *UserM) ToStringAll() string {
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
	db.Migrator().DropTable(&UserM{})
	db.Migrator().CreateTable(&UserM{})
	db.Migrator().DropTable(&CreditCard{})
	db.Migrator().CreateTable(&CreditCard{})
	println("Table created")

}

func initDataUserHasMany() {
	//.Debug()
	db.Save(&UserM{
		Name: "Mostafa",
		CreditCards: []CreditCard{
			{Number: "111"},
			{Number: "333"},
		},
	})

	db.Save(&UserM{
		Name: "Hasan",
		CreditCards: []CreditCard{
			{Number: "222"},
		},
	})
}

func PreloadUserHasMany() {

	users := []UserM{}

	// db.Preload("CreditCards").Find(&users)
	if err := db.Where("id = ? and \"Name\" = ?", 1, "Mostafa").Preload("CreditCards").Find(&users).Error; err != nil {
		fmt.Printf(err.Error())
	}
	printUsers(users)
}

func JoinsUserHasMany() {

	users := []UserM{}
	tbl := `JOIN "CreditCard" ON "CreditCard"."UserID" = "UserM"."id"`

	if err := db.Debug().Joins(tbl).Find(&users).Error; err != nil {
		fmt.Printf(err.Error())
	}
	printUsers(users)
}

func printUsers(users []UserM) {
	for _, user := range users {
		fmt.Println(user.ToStringAll())
	}

}

func RowsUserHasMany() {

	rows, err := db.Debug().Table("UserM").Where(`"UserM"."Name" = ?`, "Mostafa").Joins(`JOIN "CreditCard" ON "CreditCard"."UserID" = "UserM"."id"`).Select("*").Rows()
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer rows.Close()

	var user UserM

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
	
	FROM "UserM" u
	LEFT JOIN "CreditCard" c ON c."UserID" = u.id
	`
)

func RawSqlUserHasMany() {

	rows, err := db.Debug().Raw(rawSQL).Rows()
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer rows.Close()

	users := make([]UserM, 0)

	for rows.Next() {
		user := UserM{}
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
