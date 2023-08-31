package labels_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/config"
	entity "github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/repository/labels"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	queries, err := mysql.New("")
	assert.Nil(t, err)
	label := labels.New(queries)
	assert.NotNil(t, label)
}

func TestCRUD(t *testing.T) {
	c, err := config.New()
	assert.Nil(t, err)
	assert.NotNil(t, c)
	dataSourceName := c.GetDB().GetDataSourceName()
	queries, err := mysql.New(dataSourceName)
	assert.Nil(t, err)
	lb := labels.New(queries)
	assert.NotNil(t, lb)
	createDTO := &entity.CreateDTO{
		Customer:   "TPLINK CRUD C",
		Model:      "repo cre",
		PartNumber: "repo cre",
		Station:    "repo cre",
		Dpi:        300,
		Label:      "repo cre",
		Setup:      `[{"variable":"{{ model }}","reportID":"1708","reportName":"WRC130BZ01E_F5_HEAD","start":"","x":""},{"variable":"{{ carton }}","reportID":"1708","reportName":"WRC130BZ01E_F5_HEAD","start":"","x":""},{"variable":"{{ serials1 }}","reportID":"1706","reportName":"WRC130BZ01E_F5_SER1","start":"1060","x":"-200"},{"variable":"{{ serials2 }}","reportID":"1707","reportName":"WRC130BZ01E_F5_SER2","start":"1060","x":"-220"}]`,
		SqlQueries: `{"{{ carton }}":"1695","{{ serials }}":"1695","{{ model }}":"1695"}`,
		Author:     "bc0g8984",
	}
	created, err := lb.Create(createDTO)
	assert.Nil(t, err)
	assert.NotNil(t, created)
	updateDTO := &entity.UpdateDTO{
		ID:         created.ID,
		Customer:   "TPLINK CRUD U",
		Model:      "repo upt",
		PartNumber: "repo upt",
		Station:    "repo upt",
		Dpi:        300,
		Label:      "repo upt",
		Setup:      `[{"variable":"{{ model }}","reportID":"1708","reportName":"WRC130BZ01E_F5_HEAD","start":"","x":""},{"variable":"{{ carton }}","reportID":"1708","reportName":"WRC130BZ01E_F5_HEAD","start":"","x":""},{"variable":"{{ serials1 }}","reportID":"1706","reportName":"WRC130BZ01E_F5_SER1","start":"1060","x":"-200"},{"variable":"{{ serials2 }}","reportID":"1707","reportName":"WRC130BZ01E_F5_SER2","start":"1060","x":"-220"}]`,
		SqlQueries: `{"{{ carton }}":"1695","{{ serials }}":"1695","{{ model }}":"1695"}`,
		Author:     "bc0g8984",
	}
	updated, err := lb.Update(updateDTO)
	assert.Nil(t, err)
	assert.NotNil(t, updated)
	readed, err := lb.GetByID(int(updated.ID))
	assert.Nil(t, err)
	assert.NotNil(t, readed)
	err = lb.DeleteByID(int(readed.ID))
	assert.Nil(t, err)
}

func TestGetAll(t *testing.T) {
	c, err := config.New()
	assert.Nil(t, err)
	assert.NotNil(t, c)
	dataSourceName := c.GetDB().GetDataSourceName()
	queries, err := mysql.New(dataSourceName)
	assert.Nil(t, err)
	lb := labels.New(queries)
	assert.NotNil(t, lb)
	list, err := lb.ListPaginate(10, 0)
	assert.Nil(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
}
