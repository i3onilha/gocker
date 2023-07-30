package config

import (
	"github.com/spf13/viper"
)

type config struct {
	db *db
}

type db struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string
}

func New() *config {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()
	return &config{
		db: &db{
			DBDriver:   viper.GetString("DB_DRIVER"),
			DBHost:     viper.GetString("DB_HOST"),
			DBPort:     viper.GetString("DB_PORT"),
			DBDatabase: viper.GetString("DB_DATABASE"),
			DBUsername: viper.GetString("DB_USERNAME"),
			DBPassword: viper.GetString("DB_PASSWORD"),
		},
	}
}

func (c *config) GetDB() *db {
	return c.db
}
