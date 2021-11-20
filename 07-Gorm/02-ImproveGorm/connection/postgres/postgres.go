package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	HOST     string `yaml:"HOST"`
	PORT     int    `yaml:"PORT"`
	DBNAME   string `yaml:"DBNAME"`
	USER     string `yaml:"USER"`
	PASSWORD string `yaml:"PASSWORD"`
	SSLMODE  string `yaml:"SSLMODE"`
	DEBUG    bool   `yaml:"DEBUG"`
}

type Connections interface {
	OpenGORM() (*gorm.DB, error)
}

type databaseConfig struct {
	config Config
	domain string
}

func (dc *databaseConfig) connectionString() string {
	fmt.Printf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s",
		dc.config.HOST, dc.config.PORT, dc.config.USER, dc.config.PASSWORD, dc.config.DBNAME, dc.config.SSLMODE)

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s",
		dc.config.HOST, dc.config.PORT, dc.config.USER, dc.config.PASSWORD, dc.config.DBNAME, dc.config.SSLMODE)
}

// CreateConnection Creates a connection to Postgres database using config and domain name
func CreateConnection(config Config, domain string) Connections {
	return &databaseConfig{
		config: config,
		domain: domain,
	}
}

// OpenGORM creates a new *gorm.DB Database instance
func (dc *databaseConfig) OpenGORM() (*gorm.DB, error) {
	config := &gorm.Config{}

	dsn := dc.connectionString()
	database, err := gorm.Open(postgres.Open(dsn), config)

	if err != nil {
		fmt.Printf("OpenGOR::connection failed: %s", err.Error())
		return nil, err
	}
	return database, err
}
