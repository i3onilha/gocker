package labels

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/i3onilha/MESEnterpriseSmart/config"
	"github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	repository "github.com/i3onilha/MESEnterpriseSmart/internal/infra/repository/labels"
	validator "github.com/i3onilha/MESEnterpriseSmart/internal/infra/validator/labels"
	usecase "github.com/i3onilha/MESEnterpriseSmart/internal/usecase/labels"
)

type Usecase interface {
	Create(dto *labels.CreateDTO) (*labels.CreateDTO, error)
	DeleteByID(id int) error
	ListByModel(customer, model string) ([]*labels.CreateDTO, error)
	ListByParts(customer, partNumber string) ([]*labels.CreateDTO, error)
	ListByModelAndStationAndDpi(partNumber, station string, dpi int) ([]*labels.CreateDTO, error)
	ListByPartsAndStationAndDpi(partNumber, station string, dpi int) ([]*labels.CreateDTO, error)
	Update(dto *labels.UpdateDTO) (*labels.CreateDTO, error)
}

func getSQLQueriesFromSetup(session string, setup []labels.Setup) ([]byte, error) {
	SQLs := make(map[string]string)
	for _, setup := range setup {
		reportID := setup.ReportID
		data := url.Values{
			"funcao":       {"CDQCREPORTSQLT::getSQLReport"},
			"conn":         {"padb"}, // padb is the name of the connection in the report server
			"dados[d][id]": {reportID},
		}
		url := "http://10.58.64.198:8081/workshop/webservice.php?session=" + session //#TODO: change to config
		resp, err := http.PostForm(url, data)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var repSQL RepSQL
		err = json.Unmarshal(body, &repSQL)
		if err != nil {
			msg := "please check if report id " + reportID + " exists: "
			return nil, errors.New(msg + err.Error())
		}
		key := setup.ReportID + "_" + setup.ReportName
		value := repSQL.REP_QUERY
		if setup.Start != "" {
			value = strings.Replace(value, "':START'", setup.Start, -1)
		}
		if setup.X != "" {
			value = strings.Replace(value, "':X'", setup.X, -1)
		}
		SQLs[key] = value
	}
	return json.Marshal(SQLs)
}

func getUsecase() (Usecase, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}
	queries, err := mysql.New(c.GetDB().GetDataSourceName())
	if err != nil {
		return nil, err
	}
	repo := repository.New(queries)
	vali := validator.New()
	return usecase.New(repo, vali), nil
}

type RepSQL struct {
	REP_QUERY string `json:"REP_QUERY"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var createLabelDTO labels.CreateLabelDTO
	session := chi.URLParam(r, "session")
	err := json.NewDecoder(r.Body).Decode(&createLabelDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlQueries, err := getSQLQueriesFromSetup(session, createLabelDTO.Setup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createDto := &labels.CreateDTO{
		Name:       createLabelDTO.Name,
		Customer:   createLabelDTO.Customer,
		Model:      createLabelDTO.Model,
		PartNumber: createLabelDTO.PartNumber,
		Station:    createLabelDTO.Station,
		Dpi:        createLabelDTO.Dpi,
		Label:      createLabelDTO.Label,
		Setup:      createLabelDTO.Setup,
		Author:     createLabelDTO.Author,
		SqlQueries: string(sqlQueries),
	}
	usec, err := getUsecase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created, err := usec.Create(createDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := strconv.Itoa(int(created.ID))
	resp := map[string]string{
		"status":  "OK",
		"message": "Label created successfully",
		"id":      id,
	}
	buf, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(buf)
}

func ListByModel(w http.ResponseWriter, r *http.Request) {
	customer := chi.URLParam(r, "customer")
	model := chi.URLParam(r, "model")
	usec, err := getUsecase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	list, err := usec.ListByModel(customer, model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	buf, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(buf)
}

func ListByParts(w http.ResponseWriter, r *http.Request) {
	customer := chi.URLParam(r, "customer")
	partNumber := chi.URLParam(r, "part_number")
	usec, err := getUsecase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	list, err := usec.ListByParts(customer, partNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	buf, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(buf)
}

func ListByModelAndStationAndDpi(w http.ResponseWriter, r *http.Request) {
	model := chi.URLParam(r, "model")
	station := chi.URLParam(r, "station")
	dpi := chi.URLParam(r, "dpi")
	dpiNumber, err := strconv.Atoi(dpi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usec, err := getUsecase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	list, err := usec.ListByModelAndStationAndDpi(model, station, dpiNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	buf, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(buf)
}

func ListByPartsAndStationAndDpi(w http.ResponseWriter, r *http.Request) {
	partNumber := chi.URLParam(r, "part_number")
	station := chi.URLParam(r, "station")
	dpi := chi.URLParam(r, "dpi")
	dpiNumber, err := strconv.Atoi(dpi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usec, err := getUsecase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	list, err := usec.ListByPartsAndStationAndDpi(partNumber, station, dpiNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	buf, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(buf)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var updateLabelDTO labels.UpdateLabelDTO
	session := chi.URLParam(r, "session")
	err := json.NewDecoder(r.Body).Decode(&updateLabelDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(updateLabelDTO.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlQueries, err := getSQLQueriesFromSetup(session, updateLabelDTO.Setup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateDto := &labels.UpdateDTO{
		ID:         int32(id),
		Name:       updateLabelDTO.Name,
		Customer:   updateLabelDTO.Customer,
		Model:      updateLabelDTO.Model,
		PartNumber: updateLabelDTO.PartNumber,
		Station:    updateLabelDTO.Station,
		Dpi:        updateLabelDTO.Dpi,
		Label:      updateLabelDTO.Label,
		Setup:      updateLabelDTO.Setup,
		Author:     updateLabelDTO.Author,
		SqlQueries: string(sqlQueries),
	}
	usec, err := getUsecase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = usec.Update(updateDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := map[string]string{
		"status":  "OK",
		"message": "Label updated successfully",
		"id":      updateLabelDTO.ID,
	}
	buf, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(buf)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	usec, err := getUsecase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = usec.DeleteByID(idNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := map[string]string{
		"status":  "OK",
		"message": "Label deleted successfully",
		"id":      id,
	}
	buf, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(buf)
}
