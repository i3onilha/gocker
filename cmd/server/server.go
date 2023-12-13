package main

import (
	"fmt"
	"net/http"

	"github.com/i3onilha/mes-spi-server/internal/router"
	"github.com/joho/godotenv"
)

var port string

func init() {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	}
	port = envFile["PORT"]
}

func main() {
	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(port, router.Router())
}
