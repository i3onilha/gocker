package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"github.com/i3onilha/MESEnterpriseSmart/config"
	"github.com/i3onilha/MESEnterpriseSmart/internal/entity/labels"
	"github.com/i3onilha/MESEnterpriseSmart/internal/infra/mysql"
	repository "github.com/i3onilha/MESEnterpriseSmart/internal/infra/repository/labels"
	validator "github.com/i3onilha/MESEnterpriseSmart/internal/infra/validator/labels"
	usecase "github.com/i3onilha/MESEnterpriseSmart/internal/usecase/labels"
)

type RepSQL struct {
	REP_QUERY string `json:"REP_QUERY"`
}

type RepLabel struct {
	Label   string  `json:"label"`
	Queries []Query `json:"queries"`
}

type Query struct {
	RepID   string   `json:"rep_id"`
	RepName string   `json:"rep_name"`
	LoopVar bool     `json:"loop_var"`
	Columns []string `json:"columns"`
	SQL     string   `json:"sql"`
}

func main() {
	log.Println("Starting server on port 7192")
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"status": "OK"}`))
		})
		r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("v0.0.1"))
		})
		r.Route("/labels", func(r chi.Router) {
			r.Post("/{session}", func(w http.ResponseWriter, r *http.Request) {
				var createLabelDTO labels.CreateLabelDTO
				session := chi.URLParam(r, "session")
				err := json.NewDecoder(r.Body).Decode(&createLabelDTO)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				// extract this to a function start
				SQLs := make(map[string]string)
				for _, setup := range createLabelDTO.Setup {
					reportID := setup.ReportID
					data := url.Values{
						"funcao":       {"CDQCREPORTSQLT::getSQLReport"},
						"conn":         {"padb"},
						"dados[d][id]": {reportID},
					}
					url := "http://10.58.64.198:8081/workshop/webservice.php?session=" + session //#TODO: change to config
					resp, err := http.PostForm(url, data)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					var repSQL RepSQL
					err = json.Unmarshal(body, &repSQL)
					if err != nil {
						msg := "please check if report id " + reportID + " exists: "
						http.Error(w, msg+err.Error(), http.StatusBadRequest)
						return
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
				sqlQueries, err := json.Marshal(SQLs)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				createDto := &labels.CreateDTO{
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
				created, err := usec.Create(createDto)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				// extract this to a function and
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
			})
			r.Get("/{part_number}", func(w http.ResponseWriter, r *http.Request) {
				partNumber := chi.URLParam(r, "part_number")
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
				list, err := usec.ListByParts(partNumber)
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
			})
			r.Get("/{part_number}/{station}/{dpi}", func(w http.ResponseWriter, r *http.Request) {
				partNumber := chi.URLParam(r, "part_number")
				station := chi.URLParam(r, "station")
				dpi := chi.URLParam(r, "dpi")
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
			})
			r.Put("/{session}", func(w http.ResponseWriter, r *http.Request) {
				var updateLabelDTO labels.UpdateLabelDTO
				session := chi.URLParam(r, "session")
				err := json.NewDecoder(r.Body).Decode(&updateLabelDTO)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				// extract this to a function start
				id, err := strconv.Atoi(updateLabelDTO.ID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				SQLs := make(map[string]string)
				for _, setup := range updateLabelDTO.Setup {
					reportID := setup.ReportID
					data := url.Values{
						"funcao":       {"CDQCREPORTSQLT::getSQLReport"},
						"conn":         {"padb"},
						"dados[d][id]": {reportID},
					}
					url := "http://10.58.64.198:8081/workshop/webservice.php?session=" + session //#TODO: change to config
					resp, err := http.PostForm(url, data)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					var repSQL RepSQL
					err = json.Unmarshal(body, &repSQL)
					if err != nil {
						msg := "please check if report id " + reportID + " exists: "
						http.Error(w, msg+err.Error(), http.StatusBadRequest)
						return
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
				sqlQueries, err := json.Marshal(SQLs)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				updateDto := &labels.UpdateDTO{
					ID:         int32(id),
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
				_, err = usec.Update(updateDto)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				// extract this to a function and
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
			})
			r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
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
			})
		})
		r.Route("/zpl", func(r chi.Router) {
			r.Get("/{part_number}/{station}/{dpi}/{serial}/{key}", func(w http.ResponseWriter, r *http.Request) {
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
					for key, sql := range sqlQueries {
						var repId, repName = "", ""
						var loopVar bool
						columns := []string{}
						for _, set := range label.Setup {
							if set.ReportID+"_"+set.ReportName == key {
								columns = append(columns, set.Variable)
								repId = set.ReportID
								repName = set.ReportName
								loopVar = set.LoopVar
							}
						}
						sql = strings.ReplaceAll(sql, fmt.Sprintf(":%s", keyReplace), chi.URLParam(r, "serial"))
						query := Query{
							RepID:   repId,
							RepName: repName,
							LoopVar: loopVar,
							Columns: columns,
							SQL:     sql,
						}
						repLabel.Queries = append(repLabel.Queries, query)
					}
					repLabels = append(repLabels, repLabel)
				}
				json.NewEncoder(w).Encode(repLabels)
			})
		})
	})
	err := http.ListenAndServe(":7192", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
