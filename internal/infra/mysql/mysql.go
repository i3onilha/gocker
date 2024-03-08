package mysql

import (
	"database/sql"

	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql/importserials"
)

const (
	driver = "mysql"
)

type MySQL struct {
	ImportSerials  *importserials.Queries
	db             *sql.DB
	driver         string
	dataSourceName string
}

func New(dataSourceName string) (*MySQL, error) {
	conn, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	importserials := importserials.New(conn)
	return &MySQL{
		ImportSerials:  importserials,
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
