package config_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/i3onilha/MESEnterpriseSmart/config"
)

func TestNew(t *testing.T) {
	c, err := config.New()
	assert.Nil(t, err)
	db := c.GetDB()
	assert.Equal(t, "mysql", db.Driver)
	assert.Equal(t, "mysql-dev", db.Host)
	assert.Equal(t, "3306", db.Port)
	assert.Equal(t, "dbdev", db.Database)
	assert.Equal(t, "default", db.Username)
	assert.Equal(t, "secret", db.Password)
}
