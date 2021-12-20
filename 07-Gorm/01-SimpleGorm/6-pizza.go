package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Workplace struct {
	gorm.Model
	Name    string  `gorm:"size:50;not null"`
	Address string  `gorm:"size:255;not null"`
	Phone   *string `gorm:"size:20"`
	//HasMany
	Workers []*Worker `gorm:"foreignKey:WorkplaceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Worker struct {
	gorm.Model
	Name  string  `gorm:"size:61;not null"`
	Phone *string `gorm:"size:20"`

	WorkplaceID uint `gorm:"not null"`
	//Workplace   *Workplace  //if gorm :foreignKey not defined use this line

	Recipes []*Recipe `gorm:"many2many:worker_recipes"` //`gorm:"many2many:worker_recipes;foreignKey:Phone;joinForeignKey:worker_id"`
}

type Recipe struct {
	gorm.Model
	Name    string    `gorm:"size:50;not null"`
	Workers []*Worker `gorm:"many2many:worker_recipes"` //`gorm:"many2many:worker_recipes;foreignKey:ID;joinForeignKey:recipe_id"`
}

type Size struct {
	gorm.Model
	Name string `gorm:"size:20;not null"`
}

type Pizza struct {
	gorm.Model

	//Belongs To
	RecipeID uint `gorm:"not null"`
	Recipe   Recipe

	//Belongs To
	SizeID uint `gorm:"not null"`
	Size   Size
}

func BoxString(x string) *string {
	return &x
}

func Migrate(db *gorm.DB) {
	workplacePrototype := &Workplace{}
	workerPrototype := &Worker{}
	recipePrototype := &Recipe{}
	sizePrototype := &Size{}
	pizzaPrototype := &Pizza{}

	db.Migrator().DropTable(&Workplace{})
	db.Migrator().DropTable(&Worker{})
	db.Migrator().DropTable(&Recipe{})
	db.Migrator().DropTable(&Size{})
	db.Migrator().DropTable(&Pizza{})
	db.Migrator().DropTable("worker_recipes")

	db.AutoMigrate(workplacePrototype, workerPrototype, recipePrototype, sizePrototype, pizzaPrototype)

	// db.Migrator().CreateConstraint(workerPrototype, "workplaces")
	//db.Model(workerPrototype).AddForeignKey("workplace_id", "workplaces(id)", "RESTRICT", "CASCADE") //gormV1

	// db.Migrator().CreateConstraint("worker_recipes", "workers")
	//db.Table("worker_recipes").AddForeignKey("worker_id", "workers(id)", "RESTRICT", "CASCADE")//gormV1

	// db.Migrator().CreateConstraint("worker_recipes", "recipes")
	//db.Table("worker_recipes").AddForeignKey("recipe_id", "recipes(id)", "RESTRICT", "CASCADE")//gormV1
}

func Seed(db *gorm.DB) {
	recipe1 := &Recipe{
		Name: "Mozzarella",
	}
	recipe2 := &Recipe{
		Name: "Onions",
	}
	recipe3 := &Recipe{
		Name: "Napolitan",
	}
	recipe4 := &Recipe{
		Name: "Pepperoni",
	}
	db.Save(recipe1)
	db.Save(recipe2)
	db.Save(&Size{
		Name: "Personal",
	})
	db.Save(&Size{
		Name: "Small",
	})
	db.Save(&Size{
		Name: "Medium",
	})
	db.Save(&Size{
		Name: "Big",
	})
	db.Save(&Size{
		Name: "Extra Big",
	})
	workplace1 := &Workplace{
		Name:    "Workplace One",
		Address: "Fake st. 123rd",
	}
	workplace2 := &Workplace{
		Name:    "Workplace Two",
		Address: "Evergreen Terrace 742nd",
		Phone:   BoxString("(56) 123-4789"),
		Workers: []*Worker{
			{
				Name: "Mauricio Macri",
				Recipes: []*Recipe{
					recipe1, recipe2, recipe3,
				},
			},
			{
				Name: "Donald Trump",
				Recipes: []*Recipe{
					recipe1, recipe2, recipe4,
				},
			},
		},
	}
	db.Save(workplace1)
	db.Save(workplace2)
	fmt.Printf("Workplaces created:\n%v\n%v\n", workplace1, workplace2)
	fmt.Printf("Recipes created:\n%v\n", []*Recipe{recipe1, recipe2, recipe3, recipe4})
}

func ListEverything(db *gorm.DB) {
	workplaces := []Workplace{}
	db.Preload("Workers").Preload("Workers.Recipes").Find(&workplaces)
	for _, workplace := range workplaces {
		fmt.Printf("Workplace data: %v\n", workplace)
		for _, worker := range workplace.Workers {
			fmt.Printf("Worker data: %v\n", worker)
			for _, recipe := range worker.Recipes {
				fmt.Printf("Recipe data: %v\n", recipe)
			}
		}
	}
}

func ClearEverything(db *gorm.DB) {
	err1 := db.Delete(&Workplace{}).Error
	fmt.Printf("Deleting the records:\n%v\n", err1)
}
