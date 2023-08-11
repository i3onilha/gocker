package labels_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/config"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql/labels"
	"github.com/stretchr/testify/assert"
)

func TestLabelCRUD(t *testing.T) {
	ctx := context.Background()
	c, err := config.New()
	assert.Nil(t, err)
	db := c.GetDB()
	conn, err := sql.Open(db.GetDriver(), db.GetDataSourceName())
	assert.Nil(t, err)
	defer conn.Close()
	lb := labels.New(conn)
	result, err := lb.Create(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	rows, err := result.RowsAffected()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), rows)
	id, err := result.LastInsertId()
	assert.Nil(t, err)
	assert.NotEqual(t, int64(0), id)
	updateParams := labels.UpdateParams{
		ID:          int32(id),
		Customer:    "Customer",
		Family:      "Family",
		Model:       "Model",
		PartNumber:  "PartNumber",
		OrderNumber: "45D91234",
		Line:        "01B",
		Station:     "Station",
		Label:       "Label",
		Author:      "bc0g8100",
	}
	result, err = lb.Update(ctx, updateParams)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	rows, err = result.RowsAffected()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), rows)
	id2, err := result.LastInsertId()
	assert.Nil(t, err)
	assert.NotEqual(t, id, id2)
	updatedLabel, err := lb.GetByID(ctx, int32(id))
	assert.Nil(t, err)
	assert.NotNil(t, updatedLabel)
	assert.Equal(t, updateParams.Customer, updatedLabel.Customer)
	assert.Equal(t, updateParams.Family, updatedLabel.Family)
	assert.Equal(t, updateParams.Model, updatedLabel.Model)
	assert.Equal(t, updateParams.PartNumber, updatedLabel.PartNumber)
	assert.Equal(t, updateParams.OrderNumber, updatedLabel.OrderNumber)
	assert.Equal(t, updateParams.Line, updatedLabel.Line)
	assert.Equal(t, updateParams.Station, updatedLabel.Station)
	assert.Equal(t, updateParams.Label, updatedLabel.Label)
	err = lb.DeleteByID(ctx, int32(id))
	assert.Nil(t, err)
	deletedLabel, err := lb.GetByID(ctx, int32(id))
	assert.NotNil(t, err)
	assert.NotNil(t, deletedLabel)
}

func TestList(t *testing.T) {
	ctx := context.Background()
	c, err := config.New()
	assert.Nil(t, err)
	db := c.GetDB()
	conn, err := sql.Open(db.GetDriver(), db.GetDataSourceName())
	assert.Nil(t, err)
	defer conn.Close()
	lb := labels.New(conn)
	paramns := labels.ListPaginateParams{

		Limit:  10,
		Offset: 0,
	}
	result, err := lb.ListPaginate(ctx, paramns)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
