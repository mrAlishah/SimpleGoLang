package main

import (
	"log"
	"multiDB/repository"
)

func main() {
	//handler, err := repository.GetDatabaseHandler(repository.MYSQL, "gouser:gouser@/Test")
	handler, err := repository.GetDatabaseHandler(repository.SQLITE, "DB/Test.db")
	//handler, err := repository.GetDatabaseHandler(repository.POSTGRESQL, "user=postgres dbname=Test password=alish sslmode=disable")
	//handler, err := repository.GetDatabaseHandler(repository.MONGODB, "mongodb://127.0.0.1")

	if err != nil {
		log.Fatal(err)
	}
	// err = handler.Add(entity.Video{
	// 	Title:       "Gost in the body",
	// 	Description: "test",
	// 	Url:         "www",
	// })
	// log.Println(err)

	log.Println(handler.GetAll())

}
