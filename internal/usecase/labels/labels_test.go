package usecase_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/config"
	entity "github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	repository "github.com/i3onilha/MESEnterpriseSmart/internal/infra/repository/labels"
	validator "github.com/i3onilha/MESEnterpriseSmart/internal/infra/validator/labels"
	usecase "github.com/i3onilha/MESEnterpriseSmart/internal/usecase/labels"
	"github.com/stretchr/testify/assert"
)

func TestCRUD(t *testing.T) {
	c, err := config.New()
	assert.Nil(t, err)
	queries, err := mysql.New(c.GetDB().GetDataSourceName())
	assert.Nil(t, err)
	r := repository.New(queries)
	v := validator.New()
	u := usecase.New(r, v)
	assert.NotNil(t, u)
	dto := &entity.CreateDTO{
		Customer:   "TPLINK",
		Model:      "WRC914AC",
		PartNumber: "ADP-12LW943",
		Station:    "MAC/SN",
		Dpi:        152,
		Label:      "Label 1",
		SqlQueries: `{"{{ carton }}":"1695","{{ serials }}":"1695","{{ model }}":"1695"}`,
		Author:     "BC0F4533",
	}
	create, err := u.Create(dto)
	assert.Nil(t, err)
	assert.NotNil(t, create)
	dto2 := &entity.UpdateDTO{
		ID:         create.ID,
		Customer:   "TPLINK",
		Model:      "WRC914AC",
		PartNumber: "ADP-12LW943",
		Station:    "MAC/SN",
		Dpi:        152,
		Label:      "Label 3",
		SqlQueries: `{"{{ carton }}":"1695","{{ serials }}":"1695","{{ model }}":"1695"}`,
		Author:     "AC0F4533",
	}
	update, err := u.Update(dto2)
	assert.Nil(t, err)
	assert.NotNil(t, update)
	get, err := u.GetByID(int(update.ID))
	assert.Nil(t, err)
	assert.NotNil(t, get)
	err = u.DeleteByID(int(get.ID))
	assert.Nil(t, err)
	list, err := u.List(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, list)
}
