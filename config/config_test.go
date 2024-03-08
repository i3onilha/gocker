package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/i3onilha/MESEnterpriseSmart/config"
)

func TestNew(t *testing.T) {
	c, err := config.New()
	assert.Nil(t, err)
	db := c.GetDB()
	assert.Equal(t, "mysql", db.GetDriver())
	assert.Equal(t, "default:secret@tcp(mysql-dev:3306)/dbdev_sagemcom?parseTime=true", db.GetDataSourceName())
}

func TestNewError(t *testing.T) {
	os.Setenv("GOENV", "production")
	c, err := config.New()
	assert.EqualError(t, err, "check if already have the .env file: open ./.env: no such file or directory")
	assert.Nil(t, c)
}
