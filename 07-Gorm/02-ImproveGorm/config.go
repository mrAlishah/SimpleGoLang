package main

import (
	"os"
	"simplegorm/connection/postgres"

	"gopkg.in/yaml.v2"
	//"gorm.io/gorm/logger"
	//"gorm.io/driver/mysql"
)

type MainConfig struct {
	Postgres postgres.Config `yaml:"POSTGRES"`
}

// LoadConfig loads configs form provided yaml file or overrides it with env variables
func LoadConfigFile(filePath string) (*MainConfig, error) {

	cfg := MainConfig{}
	if filePath != "" {
		err := readFile(&cfg, filePath)
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func readFile(cfg *MainConfig, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}
