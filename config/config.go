package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	envFolder = "."
)

type db struct {
	Driver   string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func (d *db) GetDriver() string {
	return d.Driver
}

func (d *db) GetDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.Username, d.Password, d.Host, d.Port, d.Database)
}

type config struct {
	db *db
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
