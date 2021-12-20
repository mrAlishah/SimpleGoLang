package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type postgresCnf struct {
	HOST     string
	PORT     int16
	USER     string
	PASSWORD string
	DBNAME   string
	SSLMODE  string
}

var srvConfig = postgresCnf{
	HOST:     "localhost",
	PORT:     5432,
	USER:     "postgres",
	PASSWORD: "alish",
	DBNAME:   "Test",
	SSLMODE:  "disable",
}

func GormOpen() error {
	dsn := connectionStringPostgres()
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = DB
	return err

}

func connectionStringPostgres() string {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s",
		srvConfig.HOST, srvConfig.PORT, srvConfig.USER, srvConfig.PASSWORD, srvConfig.DBNAME, srvConfig.SSLMODE)
}
