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
	model := chi.URLParam(r, "model")
	station := chi.URLParam(r, "station")
	dpi := chi.URLParam(r, "dpi")
	keyReplace := chi.URLParam(r, "key")
	dpiNumber, err := strconv.Atoi(dpi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c, err := config.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	queries, err := mysql.New(c.GetDB().GetDataSourceName())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repo := repository.New(queries)
	vali := validator.New()
	usec := usecase.New(repo, vali)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	list, err := usec.ListZPLByModelAndStationAndDpi(model, station, dpiNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
			http.Error(w, err.Error(), http.StatusBadRequest)
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
			d, err := execQuery(sqlQuery, keyReplace, chi.URLParam(r, "serial"), loopVar)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
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
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			repLabel.Data = append(repLabel.Data, data)
		}
		repLabels = append(repLabels, repLabel)
	}
	json.NewEncoder(w).Encode(repLabels)
}

func GetZPLCodeByPartnumber(w http.ResponseWriter, r *http.Request) {
	partNumber := chi.URLParam(r, "part_number")
	station := chi.URLParam(r, "station")
	dpi := chi.URLParam(r, "dpi")
	keyReplace := chi.URLParam(r, "key")
	dpiNumber, err := strconv.Atoi(dpi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c, err := config.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	queries, err := mysql.New(c.GetDB().GetDataSourceName())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repo := repository.New(queries)
	vali := validator.New()
	usec := usecase.New(repo, vali)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	list, err := usec.ListZPLByPartsAndStationAndDpi(partNumber, station, dpiNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
			http.Error(w, err.Error(), http.StatusBadRequest)
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
			d, err := execQuery(sqlQuery, keyReplace, chi.URLParam(r, "serial"), loopVar)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
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
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			repLabel.Data = append(repLabel.Data, data)
		}
		repLabels = append(repLabels, repLabel)
	}
	json.NewEncoder(w).Encode(repLabels)
}

func execQuery(sqlQuery, key, value string, loopVar bool) (string, error) {
	sqlQuery = strings.ReplaceAll(sqlQuery, fmt.Sprintf(":%s", key), value)
	// extract to config start
	db, err := sql.Open("godror", `user="tmcp" password="padboratmcp" connectString="10.57.64.131:1521/PADB"`)
	if err != nil {
		return "", err
	}
	defer db.Close()
	// extract to config end
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
	}
	jsonStr = strings.TrimSuffix(jsonStr, ",")
	return jsonStr, nil
}
