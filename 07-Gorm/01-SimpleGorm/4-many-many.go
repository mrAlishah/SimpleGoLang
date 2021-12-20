package main

import (
	"fmt"

	"gorm.io/gorm"
)

// User has and belongs to many languages, use `user_languages` as join table
type UserM2M struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages;foreignKey:ID;joinForeignKey:user_id;References:ID"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*UserM2M `gorm:"many2many:user_languages;foreignKey:ID;joinForeignKey:lang_id;References:ID"`
}

type FirstTbl struct {
	gorm.Model
	Name string
}

//:: Other
func initSchemaUserMany2Many() {
	// Migrate the schema
	db.Migrator().DropTable(&FirstTbl{})
	db.Migrator().DropTable(&UserM2M{})
	db.Migrator().DropTable(&Language{})
	db.Migrator().DropTable("user_languages")

	db.Debug().AutoMigrate(&FirstTbl{}, &Language{}, &UserM2M{})

	println("Table created")

}

func initDataUserMany2Many() {

	lang1 := &Language{Name: "En"}
	lang2 := &Language{Name: "Fa"}
	lang3 := &Language{Name: "NL"}
	lang4 := &Language{Name: "De"}

	db.Save(lang1)
	db.Save(lang2)
	db.Save(lang3)
	db.Save(lang4)

	db.Save(&UserM2M{
		Name: "ali",
		Languages: []*Language{
			{Name: "En"},
			{Name: "Fa"},
		},
	})

	db.Save(&UserM2M{
		Name: "Joo",
		Languages: []*Language{
			lang3,
			lang4,
		},
	})
}

func PreloadUserMany2Many() {

	// Query associated users from Language
	fmt.Println("PreloadUserMany2Many")
	users := []UserM2M{}
	u := UserM2M{Name: "ali"}
	db.Debug().Where(&u).Preload("Languages").Find(&users)

	printUserM2M(users)

}

func printUserM2M(users []UserM2M) {

	for _, user := range users {
		fmt.Printf("user : %v\n", user)
		for _, lang := range user.Languages {
			fmt.Printf("lang : %v\n", lang)

		}
	}
}
