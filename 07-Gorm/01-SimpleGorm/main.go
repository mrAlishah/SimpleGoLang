package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresCnf struct {
	HOST     string
	PORT     int16
	USER     string
	PASSWORD string
	DBNAME   string
	SSLMODE  string
}

func main() {
	srvConfig := postgresCnf{
		HOST:     "localhost",
		PORT:     5432,
		USER:     "postgres",
		PASSWORD: "alish",
		DBNAME:   "postgres",
		SSLMODE:  "disable",
	}
	dsn := connectionStringPostgres(srvConfig)
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	println("Connection to database established")
}

func connectionStringPostgres(config postgresCnf) string {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s",
		config.HOST, config.PORT, config.USER, config.PASSWORD, config.DBNAME, config.SSLMODE)
}
