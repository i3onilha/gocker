package mysql

import (
	"database/sql"

	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql/labels"
)

const (
	driver = "mysql"
)

type MySQL struct {
	driver string
	Labels *labels.Queries
}

func New(dataSourceName string) (*MySQL, error) {
	conn, err := sql.Open(labels.Driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	labels := labels.New(conn)
	return &MySQL{
		Labels: labels,
		driver: driver,
	}, nil
}

func (m *MySQL) GetDriver() string {
	return m.driver
}
