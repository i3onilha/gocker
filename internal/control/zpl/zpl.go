package zpl

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror" // v0.35.1
	"github.com/i3onilha/MESEnterpriseSmart/config"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	repository "github.com/i3onilha/MESEnterpriseSmart/internal/infra/repository/labels"
	validator "github.com/i3onilha/MESEnterpriseSmart/internal/infra/validator/labels"
	usecase "github.com/i3onilha/MESEnterpriseSmart/internal/usecase/labels"
)

type RepLabel struct {
	Label string                   `json:"label"`
	Data  []map[string]interface{} `json:"data"`
}

func GetZPLCodeByModel(w http.ResponseWriter, r *http.Request) {
	customer := chi.URLParam(r, "customer")
	model := chi.URLParam(r, "model")
	station := strings.ReplaceAll(chi.URLParam(r, "station"), "%2F", "/")
	dpi := chi.URLParam(r, "dpi")
	keyReplace := chi.URLParam(r, "key")
	dpiNumber, err := strconv.Atoi(dpi)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	c, err := config.New()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	queries, err := mysql.New(c.GetDB().GetDataSourceName())
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	defer queries.Close()
	repo := repository.New(queries)
	vali := validator.New()
	usec := usecase.New(repo, vali)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	list, err := usec.ListZPLByModelAndStationAndDpi(customer, model, station, dpiNumber)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Get ZPL List %s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	oracleDataSource, err := usec.GetOracleDataSource(customer)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Oracle data source %s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	repLabels := []RepLabel{}
	for _, label := range list {
		repLabel := RepLabel{
			Label: label.Label,
		}
		sqlQueries := make(map[string]string)
		err := json.Unmarshal([]byte(label.SqlQueries), &sqlQueries)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "JSON encoding SQL %s"}`, err.Error()), http.StatusBadRequest)
			return
		}
		for key, sqlQuery := range sqlQueries {
			var loopVar bool
			columns := []string{}
			for _, set := range label.Setup {
				if set.ReportID+"_"+set.ReportName == key {
					columns = append(columns, set.Variable)
					loopVar = set.LoopVar
				}
			}
			d, err := execQuery(oracleDataSource, sqlQuery, keyReplace, chi.URLParam(r, "serial"), loopVar)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"error": "%s FROM (%s) LABEL"}`, err.Error(), label.Name), http.StatusBadRequest)
				return
			}
			if loopVar {
				column := strings.Join(columns, "")
				column = strings.TrimPrefix(column, "{{ ")
				column = strings.TrimSuffix(column, " }}")
				column = strings.TrimPrefix(column, "{{")
				column = strings.TrimSuffix(column, "}}")
				d = fmt.Sprintf(`{"%s":%s}`, column, d)
			}
			var data map[string]interface{}
			err = json.Unmarshal([]byte(d), &data)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"error": "JSON encoding data %s from %s"}`, err.Error(), label.Name), http.StatusBadRequest)
				return
			}
			repLabel.Data = append(repLabel.Data, data)
		}
		repLabels = append(repLabels, repLabel)
	}
	json.NewEncoder(w).Encode(repLabels)
}

func GetZPLCodeByPartnumber(w http.ResponseWriter, r *http.Request) {
	customer := chi.URLParam(r, "customer")
	partNumber := chi.URLParam(r, "part_number")
	station := strings.ReplaceAll(chi.URLParam(r, "station"), "%2F", "/")
	dpi := chi.URLParam(r, "dpi")
	keyReplace := chi.URLParam(r, "key")
	dpiNumber, err := strconv.Atoi(dpi)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	c, err := config.New()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	queries, err := mysql.New(c.GetDB().GetDataSourceName())
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	defer queries.Close()
	repo := repository.New(queries)
	vali := validator.New()
	usec := usecase.New(repo, vali)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	list, err := usec.ListZPLByPartsAndStationAndDpi(customer, partNumber, station, dpiNumber)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	oracleDataSource, err := usec.GetOracleDataSource(customer)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	repLabels := []RepLabel{}
	for _, label := range list {
		repLabel := RepLabel{
			Label: label.Label,
		}
		sqlQueries := make(map[string]string)
		err := json.Unmarshal([]byte(label.SqlQueries), &sqlQueries)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
			return
		}
		for key, sqlQuery := range sqlQueries {
			var loopVar bool
			columns := []string{}
			for _, set := range label.Setup {
				if set.ReportID+"_"+set.ReportName == key {
					columns = append(columns, set.Variable)
					loopVar = set.LoopVar
				}
			}
			d, err := execQuery(oracleDataSource, sqlQuery, keyReplace, chi.URLParam(r, "serial"), loopVar)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
				return
			}
			if loopVar {
				column := strings.Join(columns, "")
				column = strings.TrimPrefix(column, "{{ ")
				column = strings.TrimSuffix(column, " }}")
				column = strings.TrimPrefix(column, "{{")
				column = strings.TrimSuffix(column, "}}")
				d = fmt.Sprintf(`{"%s":%s}`, column, d)
			}
			var data map[string]interface{}
			err = json.Unmarshal([]byte(d), &data)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
				return
			}
			repLabel.Data = append(repLabel.Data, data)
		}
		repLabels = append(repLabels, repLabel)
	}
	json.NewEncoder(w).Encode(repLabels)
}

func execQuery(oracleDataSource, sqlQuery, key, value string, loopVar bool) (string, error) {
	sqlQuery = strings.ReplaceAll(sqlQuery, fmt.Sprintf(":%s", key), value)
	db, err := sql.Open("godror", oracleDataSource)
	if err != nil {
		return "", err
	}
	defer db.Close()
	rows, err := db.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return "", err
	}
	jsonStr, err := jsonSerialize(rows)
	if err != nil {
		return "", err
	}
	if loopVar {
		jsonStr = fmt.Sprintf("[%s]", jsonStr)
	}
	return jsonStr, nil
}

func jsonSerialize(rows *sql.Rows) (string, error) {
	colNames, err := rows.Columns()
	if err != nil {
		return "", err
	}
	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}
	var result = make(map[string]interface{})
	var jsonStr string
	var count int
	for rows.Next() {
		err = rows.Scan(colPtrs...)
		if err != nil {
			return "", err
		}
		for i, col := range cols {
			result[colNames[i]] = col
		}
		jsonStr += "{"
		for key, val := range result {
			jsonStr += fmt.Sprintf(`"%s":"%s",`, key, val)
		}
		jsonStr = strings.TrimSuffix(jsonStr, ",")
		jsonStr += "},"
		count++
	}
	if count == 0 {
		return "", fmt.Errorf("NO DATA FOUND TO (%s) COLUMNS", strings.Join(colNames, ", "))
	}
	jsonStr = strings.TrimSuffix(jsonStr, ",")
	return jsonStr, nil
}
