package labels

import (
	"context"
	"database/sql"
)

const (
	Driver = "mysql"
)

type CreateParams struct {
	Customer   string
	Family     string
	Model      string
	PartNumber string
	Station    string
	Label      string
	Author     string
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
	result, err = queries.Create(ctx)
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
		Family:     data.Family,
		Model:      data.Model,
		PartNumber: data.PartNumber,
		Station:    data.Station,
		Label:      data.Label,
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
