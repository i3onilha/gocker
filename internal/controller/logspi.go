package controller

import (
	"encoding/json"
	"net/http"

	"github.com/i3onilha/mes-spi-server/internal/model"
)

type Resp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Write struct {
	w       http.ResponseWriter
	message string
	err     error
}

func SPILogger(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var data model.DataLog
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		httpWrite(Write{
			w:       w,
			message: err.Error(),
			err:     err,
		})
		return
	}
	message, err := model.LogInDatabase(&data)
	httpWrite(Write{
		w:       w,
		message: message,
		err:     err,
	})
}

func httpWrite(write Write) {
	header := http.StatusOK
	status := "OK"
	if write.err != nil {
		header = http.StatusInternalServerError
		status = "NOK"
	}
	write.w.WriteHeader(header)
	resp := Resp{
		Status:  status,
		Message: write.message,
	}

	json.NewEncoder(write.w).Encode(resp)
}
