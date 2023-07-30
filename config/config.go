package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	envFolder = "."
)

type config struct {
	db *db
}

type db struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func New() (*config, error) {
	if os.Getenv("GOENV") == "development" {
		envFolder = ".."
	}
	path := fmt.Sprintf("%s/.env", envFolder)
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("check if already have the .env file: %v", err)
	}
	return &config{
		db: &db{
			Driver:   viper.GetString("DB_DRIVER"),
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			Database: viper.GetString("DB_DATABASE"),
			Username: viper.GetString("DB_USERNAME"),
			Password: viper.GetString("DB_PASSWORD"),
		},
	}, nil
}

func (c *config) GetDB() *db {
	return c.db
}
