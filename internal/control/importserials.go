package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/unmarshalcsv"
	"github.com/i3onilha/MESEnterpriseSmart/internal/service/importpallet"
)

type Res struct {
	Status  string `json:"status"`
	Type    string `json:"type"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

func GetDataCSVFile(w http.ResponseWriter, r *http.Request) {
	comma := []rune(chi.URLParam(r, "comma"))
	if len(comma) != 1 {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: "Comma format not allowed",
		}
		json.NewEncoder(w).Encode(res)
	}
	csvBuf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	data, err := unmarshalcsv.UnmarshalCSV(csvBuf, comma[0])
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func SaveList(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	comma := []rune(chi.URLParam(r, "comma"))
	csvBuf, err := ioutil.ReadAll(r.Body)
	ctx := r.Context()
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	importParams := importpallet.ImportParams{
		UUID:   uuid,
		Comma:  comma[0],
		CsvBuf: csvBuf,
	}
	importer, err := importpallet.NewImportPallet(ctx, importParams)
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	err = importer.ImportSerial(uuid)
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	res := Res{
		Status:  "OK",
		Type:    "info",
		Message: "SERIALS IMPORTED WITH SUCCESS.",
	}
	json.NewEncoder(w).Encode(res)
}

func GetByPallet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := chi.URLParam(r, "key")
	value := chi.URLParam(r, "value")
	importParams := importpallet.ImportParams{
		Comma: ',',
	}
	importer, err := importpallet.NewImportPallet(ctx, importParams)
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	res, err := importer.GetList(key, value)
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func CheckPallet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pallet := chi.URLParam(r, "pallet")
	importParams := importpallet.ImportParams{
		Comma: ',',
	}
	importer, err := importpallet.NewImportPallet(ctx, importParams)
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	res, err := importer.CheckPallet(pallet)
	if err != nil {
		res := Res{
			Status:  "NOK",
			Type:    "danger",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	if res == nil {
		res := Res{
			Status:  "OK",
			Type:    "info",
			Message: fmt.Sprintf("O pallet %s ainda nao foi salvo", pallet),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(res)
}
