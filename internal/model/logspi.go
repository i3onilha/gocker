package model

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

type DataLog struct {
	Model   string `json:"model"`
	Station string `json:"station"`
	File    string `json:"file"`
	Content string `json:"content"`
}

func LogInDatabase(data *DataLog) (string, error) {
	db, err := sql.Open("godror", `user="tmcp" password="PBDBORATMCP" connectString="10.57.64.131:1521/PBDB"`)
	if err != nil {
		return "HOUVE UM ERRO AO TENTAR CONECTAR COM BANCO DE DADOS.", err
	}
	defer db.Close()
	var result string
	sql := fmt.Sprintf("SELECT FGET_PROCCESS_LOG('%s', '%s', '%s', '%s') AS RESULT FROM DUAL", data.Model, data.Station, data.File, data.Content)
	err = db.QueryRow(sql).Scan(&result)
	if err != nil {
		return "HOUVE UM ERRO AO TENTAR OBTER O RESULTADO DA FUNCAO FGET_PROCCESS_LOG COM OS PARAMETROS PASSADOS.", err
	}
	return result, nil
}
