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

func (q *Queries) CreateAndUpdateLabel(ctx context.Context, dataSourceName string, data CreateLabelParams) error {
	conn, err := sql.Open(Driver, dataSourceName)
	if err != nil {
		return err
	}
	defer conn.Close()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	queries := q.WithTx(tx)
	if err != nil {
		return err
	}
	resultCreated, err := queries.CreateLabel(ctx)
	if err != nil {
		return err
	}
	id, err := resultCreated.LastInsertId()
	if err != nil {
		return err
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
	_, err = queries.UpdateLabel(ctx, params)
	if err != nil {
		return err
	}
	return tx.Commit()
}
