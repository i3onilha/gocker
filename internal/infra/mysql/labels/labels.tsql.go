package labels

import (
	"context"
	"database/sql"
)

const (
	Driver = "mysql"
)

type CreateParams struct {
	Customer    string
	Model       string
	PartNumber  string
	OrderNumber string
	Line        string
	Station     string
	Dpi         int32
	Label       string
	Setup       string
	SqlQueries  string
	Author      string
}

func (q *Queries) CreateAndUpdate(ctx context.Context, dataSourceName string, data CreateParams) (int, error) {
	conn, err := sql.Open(Driver, dataSourceName)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	tx, err := conn.Begin()
	if err != nil {
		return 0, err
	}
	queries := q.WithTx(tx)
	if err != nil {
		return 0, err
	}
	result, err := queries.Create(ctx)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	params := UpdateParams{
		ID:         int32(id),
		Customer:   data.Customer,
		Model:      data.Model,
		PartNumber: data.PartNumber,
		Station:    data.Station,
		Dpi:        data.Dpi,
		Label:      data.Label,
		Setup:      data.Setup,
		SqlQueries: data.SqlQueries,
		Author:     data.Author,
	}
	result, err = queries.Update(ctx, params)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
