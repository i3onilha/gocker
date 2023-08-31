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

func TestCreateAndUpdate(t *testing.T) {
	ctx := context.Background()
	c, err := config.New()
	assert.Nil(t, err)
	db := c.GetDB()
	conn, err := sql.Open(db.GetDriver(), db.GetDataSourceName())
	assert.Nil(t, err)
	defer conn.Close()
	lb := labels.New(conn)
	params := labels.CreateParams{
		Customer:    "TPLINK CREA",
		Setup:       `[{"variable":"{{ model }}","reportID":"1708","reportName":"WRC130BZ01E_F5_HEAD","start":"","x":""},{"variable":"{{ carton }}","reportID":"1708","reportName":"WRC130BZ01E_F5_HEAD","start":"","x":""},{"variable":"{{ serials1 }}","reportID":"1706","reportName":"WRC130BZ01E_F5_SER1","start":"1060","x":"-200"},{"variable":"{{ serials2 }}","reportID":"1707","reportName":"WRC130BZ01E_F5_SER2","start":"1060","x":"-220"}]`,
		SqlQueries:  `{"{{ carton }}":"1695","{{ serials }}":"1695","{{ model }}":"1695"}`,
		Model:       "model created",
		PartNumber:  "partnumber 2",
		OrderNumber: "38D45675",
		Line:        "03",
		Station:     "station created",
		Dpi:         300,
		Label:       "label created",
		Author:      "bc0g8101",
	}
	id, err := lb.CreateAndUpdate(ctx, db.GetDataSourceName(), params)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}
