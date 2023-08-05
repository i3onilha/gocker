package labels

import (
	"context"
	"database/sql"
)

const (
	Driver = "mysql"
)

func (q *Queries) CreateAndUpdateLabel(ctx context.Context, dataSourceName string, params UpdateLabelParams) error {
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
	params.ID = int32(id)
	_, err = queries.UpdateLabel(ctx, params)
	if err != nil {
		return err
	}
	return tx.Commit()
}
