package config_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/i3onilha/MESEnterpriseSmart/config"
)

func TestNew(t *testing.T) {
	c := config.New()
	db := c.GetDB()
	assert.Equal(t, "mysql", db.DBDriver)
	assert.Equal(t, "mysql-dev", db.DBHost)
	assert.Equal(t, "3306", db.DBPort)
	assert.Equal(t, "dbdev", db.DBDatabase)
	assert.Equal(t, "default", db.DBUsername)
	assert.Equal(t, "secret", db.DBPassword)
}
