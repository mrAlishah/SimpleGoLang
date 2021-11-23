package main

import (
	"fmt"
	pg "simplegorm/connection/postgres"
	"simplegorm/entity"
	"simplegorm/repository"
)

func main() {

	// Loading the config file
	cfg, err := LoadConfigFile("config-local.yaml")
	if err != nil {
		fmt.Printf("failed to load config: %s", err.Error())
	}

	pgConn := pg.CreateConnection(cfg.Postgres, "pgGorm")
	gorm, err := pgConn.OpenGORM()
	if err != nil {
		fmt.Printf("failed to Coonect DB: %s", err.Error())
	}
	println("Connection to database established")

	userRepo, err := repository.CreateUserRepository(gorm)
	if err != nil {
		fmt.Printf("failed to Create Repo: %s", err.Error())
	}

	userRepo.InitUserData()

	users, _ := userRepo.FindAll("")
	printAll(users)

	user, _ := userRepo.FindFirst(1)
	user.LastName = "Shokri"
	userRepo.Update(*user)

	updateUser := entity.User{
		ID:        1,
		UserName:  "666",
		FirstName: "saeed",
		LastName:  "asadolahi",
	}

	userRepo.Update(updateUser)
	users, _ = userRepo.FindAll("")
	printAll(users)

	newUser := entity.User{
		UserName:  "666",
		FirstName: "saeed",
		LastName:  "asadolahi",
	}
	userAdd, err := userRepo.Add(newUser)
	if err != nil {
		fmt.Printf("failed to Create Repo: %s", err.Error())
	}
	fmt.Printf("New Record  %v \n", userAdd)
}

func printAll(users []entity.User) {
	println("=>PrintAll:: User List")
	for _, user := range users {
		fmt.Printf("\t Record  %v \n", user)
	}
	println("==============================\n")
}
