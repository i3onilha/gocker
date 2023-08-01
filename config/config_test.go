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
	assert.Equal(t, "mysql", db.GetDriver())
	assert.Equal(t, "default:secret@tcp(mysql-dev:3306)/dbdev?parseTime=true", db.GetDataSourceName())
}
