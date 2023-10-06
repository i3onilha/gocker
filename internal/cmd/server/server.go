package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/i3onilha/MESEnterpriseSmart/internal/control/labels"
	"github.com/i3onilha/MESEnterpriseSmart/internal/control/zpl"
)

type RepSQL struct {
	REP_QUERY string `json:"REP_QUERY"`
}

type RepLabel struct {
	Label string                   `json:"label"`
	Data  []map[string]interface{} `json:"data"`
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
			r.Post("/{session}", labels.Create)
			r.Get("/{part_number}/partnumber", labels.ListByParts)
			r.Get("/{model}/model", labels.ListByModel)
			r.Get("/{model}/{station}/{dpi}", labels.ListByModelAndStationAndDpi)
			r.Get("/{part_number}/{station}/{dpi}", labels.ListByPartsAndStationAndDpi)
			r.Put("/{session}", labels.Update)
			r.Delete("/{id}", labels.Delete)
		})
		r.Route("/zpl", func(r chi.Router) {
			r.Get("/{part_number}/{station}/{dpi}/{serial}/{key}", zpl.GetZPLCode)
		})
	})
	err := http.ListenAndServe(":7192", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
