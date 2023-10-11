package copylabel

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror" // v0.35.1
	"github.com/i3onilha/MESEnterpriseSmart/config"
	entity "github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	repository "github.com/i3onilha/MESEnterpriseSmart/internal/infra/repository/labels"
	validator "github.com/i3onilha/MESEnterpriseSmart/internal/infra/validator/labels"
	usecase "github.com/i3onilha/MESEnterpriseSmart/internal/usecase/labels"
)

func CopyModel(w http.ResponseWriter, r *http.Request) {
	customer := chi.URLParam(r, "customer")
	model_from := chi.URLParam(r, "model_from")
	model_to := chi.URLParam(r, "model_to")
	station_to := chi.URLParam(r, "station_to")
	dpi := chi.URLParam(r, "dpi_to")
	dpi_to, err := strconv.Atoi(dpi)
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
	list, err := usec.ListByModel(customer, model_from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, v := range list {
		createDto := &entity.CreateDTO{
			Name:       v.Name,
			Customer:   v.Customer,
			Model:      model_to,
			PartNumber: v.PartNumber,
			Station:    station_to,
			Dpi:        int32(dpi_to),
			Label:      v.Label,
			Setup:      v.Setup,
			Author:     v.Author,
			SqlQueries: v.SqlQueries,
		}
		_, err = usec.Create(createDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Write([]byte(`{"message": "ALL LABELS COPIED SUCCESSFULLY"}`))
}
