package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/i3onilha/MESEnterpriseSmart/config"
)

type RepSQL struct {
	REP_QUERY string `json:"REP_QUERY"`
}

type RepLabel struct {
	Label string                   `json:"label"`
	Data  []map[string]interface{} `json:"data"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func main() {
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
			w.Write([]byte(`{"appname": "SAGEMCOM Service", "status": "OK"}`))
		})
		r.Route("/get-data-csv-file", func(r chi.Router) {
			r.Post("/csv", func(w http.ResponseWriter, r *http.Request) {
				csvBuf, err := ioutil.ReadAll(r.Body)
				if err != nil {
					responseError := ResponseError{
						Status:  "NOK",
						Message: err.Error(),
					}
					json.NewEncoder(w).Encode(responseError)
					return
				}
				fmt.Println(string(csvBuf))
			})
		})
	})
	var err error
	c, err := config.New()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
	port := fmt.Sprintf(":%s", c.GetPort())
	log.Println(fmt.Sprintf("Starting server on port %s", port))
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
