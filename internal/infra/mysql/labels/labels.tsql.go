package labels

import (
	"context"
	"database/sql"
)

const (
	Driver = "mysql"
)

type CreateLabelParams struct {
	Customer   string
	Family     string
	Model      string
	PartNumber string
	Station    string
	Label      string
	Author     string
}

func (q *Queries) CreateAndUpdateLabel(ctx context.Context, dataSourceName string, data CreateLabelParams) (int, error) {
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
	result, err := queries.CreateLabel(ctx)
	if err != nil {
		return 0, err
	}
	result, err = queries.CreateLabel(ctx)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	params := UpdateLabelParams{
		ID:         int32(id),
		Customer:   data.Customer,
		Family:     data.Family,
		Model:      data.Model,
		PartNumber: data.PartNumber,
		Station:    data.Station,
		Label:      data.Label,
		Author:     data.Author,
	}
	result, err = queries.UpdateLabel(ctx, params)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
