package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":1808"
	log.Printf("Running server on %s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home page"))
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
