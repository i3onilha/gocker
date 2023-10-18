package mysql

import (
	"database/sql"

	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql/labels"
)

const (
	driver = "mysql"
)

type MySQL struct {
	Labels         *labels.Queries
	db             *sql.DB
	driver         string
	dataSourceName string
}

func New(dataSourceName string) (*MySQL, error) {
	conn, err := sql.Open(labels.Driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	labels := labels.New(conn)
	return &MySQL{
		Labels:         labels,
		db:             conn,
		driver:         driver,
		dataSourceName: dataSourceName,
	}, nil
}

func (m *MySQL) Close() error {
	return m.db.Close()
}

func (m *MySQL) GetDriver() string {
	return m.driver
}

func (m *MySQL) GetDataSourceName() string {
	return m.dataSourceName
}
