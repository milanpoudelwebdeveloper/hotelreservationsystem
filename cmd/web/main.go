package main

import (
	"hotelsystem/pkg/handlers"
	"net/http"
)
const portNumber = ":8081"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.ListenAndServe(portNumber, nil)
}
