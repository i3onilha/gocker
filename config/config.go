package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type db struct {
	driver   string
	host     string
	port     string
	database string
	username string
	password string
}

func (d *db) GetDriver() string {
	return d.driver
}

func (d *db) GetDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", d.username, d.password, d.host, d.port, d.database)
}

type config struct {
	db *db
}

func New() (*config, error) {
	path := fmt.Sprintf("%s/.env", os.Getenv("SOURCE_CODE"))
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("check if already have the .env file: %v", err)
	}
	return &config{
		db: &db{
			driver:   viper.GetString("DB_DRIVER"),
			host:     viper.GetString("DB_HOST"),
			port:     viper.GetString("DB_PORT"),
			database: viper.GetString("DB_DATABASE"),
			username: viper.GetString("DB_USERNAME"),
			password: viper.GetString("DB_PASSWORD"),
		},
	}, nil
}

func (c *config) GetDB() *db {
	return c.db
}
