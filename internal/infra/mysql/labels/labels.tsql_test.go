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

func TestCreateAndUpdateLabel(t *testing.T) {
	ctx := context.Background()
	c, err := config.New()
	assert.Nil(t, err)
	db := c.GetDB()
	conn, err := sql.Open(db.GetDriver(), db.GetDataSourceName())
	assert.Nil(t, err)
	defer conn.Close()
	lb := labels.New(conn)
	params := labels.CreateLabelParams{
		Customer:   "customer created",
		Family:     "family",
		Model:      "model created",
		PartNumber: "partnumber 2",
		Station:    "station created",
		Label:      "label created",
		Author:     "bc0g8101",
	}
	id, err := lb.CreateAndUpdateLabel(ctx, db.GetDataSourceName(), params)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}
