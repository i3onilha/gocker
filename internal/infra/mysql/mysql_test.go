package mysql_test

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/config"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql/labels"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c, err := config.New()
	assert.Nil(t, err)
	dataSourceName := c.GetDB().GetDataSourceName()
	db, err := mysql.New(dataSourceName)
	assert.Nil(t, err)
	assert.NotNil(t, db)
	ctx := context.Background()
	params := labels.CreateParams{
		Customer:   "TPLINK MYSQL",
		Model:      "model created",
		PartNumber: "partnumber 2",
		Station:    "station created",
		Dpi:        300,
		Label:      "label created",
		Setup:      `[{"variable":"{{ model }}","reportID":"1708","reportName":"WRC130BZ01E_F5_HEAD","start":"","x":""},{"variable":"{{ carton }}","reportID":"1708","reportName":"WRC130BZ01E_F5_HEAD","start":"","x":""},{"variable":"{{ serials1 }}","reportID":"1706","reportName":"WRC130BZ01E_F5_SER1","start":"1060","x":"-200"},{"variable":"{{ serials2 }}","reportID":"1707","reportName":"WRC130BZ01E_F5_SER2","start":"1060","x":"-220"}]`,
		SqlQueries: `{"{{ carton }}":"1695","{{ serials }}":"1695","{{ model }}":"1695"}`,
		Author:     "bc0g8101",
	}
	id, err := db.Labels.CreateAndUpdate(ctx, dataSourceName, params)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}

func TestGetDriver(t *testing.T) {
	db, err := mysql.New("")
	assert.Nil(t, err)
	assert.NotNil(t, db)
	driver := db.GetDriver()
	assert.Equal(t, "mysql", driver)
}

func TestGetDataSourceName(t *testing.T) {
	db, err := mysql.New("default:secret@tcp(mysql-dev:3306)/dbdev?parseTime=true")
	assert.Nil(t, err)
	assert.NotNil(t, db)
	dataSourceName := db.GetDataSourceName()
	assert.Equal(t, "default:secret@tcp(mysql-dev:3306)/dbdev?parseTime=true", dataSourceName)
}
