package control

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/unmarshalcsv"
)

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func GetDataCSVFile(w http.ResponseWriter, r *http.Request) {
	comma := []rune(chi.URLParam(r, "comma"))
	csvBuf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError := Error{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(responseError)
		return
	}
	data, err := unmarshalcsv.UnmarshalCSV(csvBuf, comma[0])
	if err != nil {
		responseError := Error{
			Status:  "NOK",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(responseError)
		return
	}
	json.NewEncoder(w).Encode(data)
}
