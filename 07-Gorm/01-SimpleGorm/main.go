package main

func main() {
	err := GormOpen()
	if err != nil {
		panic(err.Error())
	}
	println("Connection to database established")

	//--------------------------------introCRUD
	//initUser()
	//introCRUD()

	//--------------------------------Relation
	//initUserHasOne()
	//initHasOne()

	//--------------------------------Has Many
	initSchemaUserHasMany()
	initDataUserHasMany()
	//PreloadUserHasMany()
	//JoinsUserHasMany()
	//RowsUserHasMany()
	RawSqlUserHasMany()
}
