package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!!"))
	})
	fmt.Println("Running webservice...")
	log.Fatal(http.ListenAndServe(":6565", nil))
}
