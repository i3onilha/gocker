package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/db/importserials"
)

type DB struct {
	ImportSerials  *importserials.Queries
	db             *sql.DB
	driver         string
	dataSourceName string
}

func New(driver, dataSourceName string) (*DB, error) {
	conn, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	importserials := importserials.New(conn)
	return &DB{
		ImportSerials:  importserials,
		db:             conn,
		driver:         driver,
		dataSourceName: dataSourceName,
	}, nil
}

func (m *DB) Close() error {
	return m.db.Close()
}

func (m *DB) GetDriver() string {
	return m.driver
}

func (m *DB) GetDataSourceName() string {
	return m.dataSourceName
}
