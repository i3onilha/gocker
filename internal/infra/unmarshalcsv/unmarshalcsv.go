package unmarshalcsv

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
)

type Row struct {
	Palete       string `json:"pallet"`
	Masterbox    string `json:"masterbox"`
	SerialNumber string `json:"serial_number"`
	Partnumber   string `json:"partnumber"`
}

func UnmarshalCSV(buf []byte, comma rune) ([]Row, error) {
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
	rows := make([]Row, 0)
	err = json.Unmarshal(csvBuf, &rows)
	if err != nil {
		return nil, err
	}
	res := make([]Row, len(rows))
	for i := 0; i < len(rows); i++ {
		res[i].Palete = rows[i].Palete
		res[i].Masterbox = rows[i].Masterbox
		res[i].SerialNumber = rows[i].SerialNumber
		res[i].Partnumber = rows[i].Partnumber
	}
	return res, nil
}
