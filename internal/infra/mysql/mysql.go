package mysql

import (
	"database/sql"

	"github.com/i3onilha/MESEnterpriseSmart/config"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql/labels"
)

type MySQL struct {
	Labels *labels.Queries
}

func New() (*MySQL, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}
	db := c.GetDB()
	conn, err := sql.Open(db.GetDriver(), db.GetDataSourceName())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	labels := labels.New(conn)
	return &MySQL{
		Labels: labels,
	}, nil
}
