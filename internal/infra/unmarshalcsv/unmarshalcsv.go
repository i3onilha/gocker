package unmarshalcsv

import (
	"bytes"
	"encoding/csv"
	"encoding/json"

	dto "github.com/i3onilha/MESEnterpriseSmart/internal/infra/dto/importpallet"
)

func UnmarshalCSV(buf []byte, comma rune) ([]dto.ResDto, error) {
	reader := csv.NewReader(bytes.NewReader(buf))
	reader.Comma = comma
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var csvMap []map[string]string
	headers := records[0]
	for i := 1; i < len(records); i++ {
		record := records[i]
		csvRow := make(map[string]string)
		for j := 0; j < len(headers) && j < len(record); j++ {
			csvRow[headers[j]] = record[j]
		}
		csvMap = append(csvMap, csvRow)
	}
	csvBuf, err := json.Marshal(csvMap)
	if err != nil {
		return nil, err
	}
	rows := make([]dto.ResDto, 0)
	err = json.Unmarshal(csvBuf, &rows)
	if err != nil {
		return nil, err
	}
	res := make([]dto.ResDto, len(rows))
	for i := 0; i < len(rows); i++ {
		res[i].Pallet = rows[i].Pallet
		res[i].MasterBox = rows[i].MasterBox
		res[i].SerialNumber = rows[i].SerialNumber
		res[i].PartNumber = rows[i].PartNumber
	}
	return res, nil
}
