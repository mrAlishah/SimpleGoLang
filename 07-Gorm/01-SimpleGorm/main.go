package main

func main() {
	err := GormOpen()
	if err != nil {
		panic(err.Error())
	}
	println("Connection to database established")

	//--------------------------------introCRUD
	initUser()
	introCRUD()
	//--------------------------------CRUD Advance
	crudAdvanced()
	//--------------------------------Relation
	// initUserHasOne()
	// initHasOne()

	//--------------------------------Has Many
	// initSchemaUserHasMany()
	// initDataUserHasMany()
	// PreloadUserHasMany()
	// JoinsUserHasMany()
	// RowsUserHasMany()
	// RawSqlUserHasMany()

	//--------------------------------Many to Many
	// initSchemaUserMany2Many()
	// initDataUserMany2Many()
	// PreloadUserMany2Many()

	//--------------------------------Pizza
	// Migrate(db)
	// Seed(db)
	// ListEverything(db)
	// ClearEverything(db)

}
