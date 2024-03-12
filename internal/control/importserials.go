package control

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/unmarshalcsv"
	"github.com/i3onilha/MESEnterpriseSmart/internal/service/importpallet"
)

type Res struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func GetDataCSVFile(w http.ResponseWriter, r *http.Request) {
	comma := []rune(chi.URLParam(r, "comma"))
	if len(comma) != 1 {
		response := Res{
			Status:  "NOK",
			Message: "Comma format not allowed",
		}
		json.NewEncoder(w).Encode(response)
	}
	csvBuf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	data, err := unmarshalcsv.UnmarshalCSV(csvBuf, comma[0])
	if err != nil {
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
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
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	importParams := importpallet.ImportParams{
		UUID:   uuid,
		Comma:  comma[0],
		CsvBuf: csvBuf,
	}
	importer, err := importpallet.NewImportPallet(ctx, importParams)
	if err != nil {
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	err = importer.ImportSerial(uuid)
	if err != nil {
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	response := Res{
		Status:  "OK",
		Message: "SERIALS IMPORTED WITH SUCCESS.",
	}
	json.NewEncoder(w).Encode(response)
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
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	response, err := importer.GetList(key, value)
	if err != nil {
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func CheckPallet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pallet := chi.URLParam(r, "pallet")
	importParams := importpallet.ImportParams{
		Comma: ',',
	}
	importer, err := importpallet.NewImportPallet(ctx, importParams)
	if err != nil {
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	response, err := importer.CheckPallet(pallet)
	if err != nil {
		response := Res{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(response)
}
