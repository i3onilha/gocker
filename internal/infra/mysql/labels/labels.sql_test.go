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

func TestCreateLabel(t *testing.T) {
	ctx := context.Background()
	c, err := config.New()
	assert.Nil(t, err)
	db := c.GetDB()
	conn, err := sql.Open(db.GetDriver(), db.GetDataSourceName())
	assert.Nil(t, err)
	labels := labels.New(conn)
	result, err := labels.CreateLabel(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	rows, err := result.RowsAffected()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), rows)
	id, err := result.LastInsertId()
	assert.Nil(t, err)
	assert.Equal(t, int64(0), id)
}
