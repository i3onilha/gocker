package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":6767"
	fmt.Println(fmt.Sprintf("Starting server on port %s", port))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := "./test/coverage.html"
		http.ServeFile(w, r, file)
	})
	log.Fatal(http.ListenAndServe(":6767", nil))
}
