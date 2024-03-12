package importpallet

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/i3onilha/MESEnterpriseSmart/internal/infra/db"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/db/importserials"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/unmarshalcsv"
)

type ImportParams struct {
	UUID   string
	Comma  rune
	CsvBuf []byte
}

type ImportPallet struct {
	ctx  context.Context
	data ImportParams
}

var runeSlice = []rune{44, 9, 32, 59}

func DoNotContainsRune(r rune) bool {
	for _, v := range runeSlice {
		if v == r {
			return false
		}
	}
	return true
}

func NewImportPallet(ctx context.Context, data ImportParams) (*ImportPallet, error) {
	if DoNotContainsRune(data.Comma) {
		return nil, fmt.Errorf("The %s separator is not allowed.", string(data.Comma))
	}
	return &ImportPallet{
		ctx:  ctx,
		data: data,
	}, nil
}

func (i *ImportPallet) ImportSerial(uuid string) error {
	data, err := unmarshalcsv.UnmarshalCSV(i.data.CsvBuf, i.data.Comma)
	if err != nil {
		return err
	}
	driver := i.ctx.Value("driver").(string)
	dataSourceName := i.ctx.Value("datasourcename").(string)
	conn, err := db.New(driver, dataSourceName)
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, item := range data {
		arg := importserials.CreateParams{
			Pallet:       sql.NullString{String: item.Pallet, Valid: true},
			Masterbox:    sql.NullString{String: item.MasterBox, Valid: true},
			SerialNumber: sql.NullString{String: item.SerialNumber, Valid: true},
			PartNumber:   sql.NullString{String: item.PartNumber, Valid: true},
			Uuid:         sql.NullString{String: uuid, Valid: true},
		}
		_, err = conn.ImportSerials.Create(i.ctx, arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *ImportPallet) GetByPallet(pallet string) ([]importserials.ImportPalletsSerial, error) {
	driver := i.ctx.Value("driver").(string)
	dataSourceName := i.ctx.Value("datasourcename").(string)
	conn, err := db.New(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.ImportSerials.GetByPallet(i.ctx, sql.NullString{
		String: pallet,
		Valid:  true,
	})
}

func (i *ImportPallet) GetList(key, value string) ([]importserials.ImportPalletsSerial, error) {
	if key == "pallet" {
		return i.GetByPallet(value)
	}
	return nil, fmt.Errorf("%s column do not exists", key)
}
