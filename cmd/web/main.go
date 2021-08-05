package main

import (
	"hotelsystem/pkg/config"
	"hotelsystem/pkg/handlers"
	"hotelsystem/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create templatecache", err)
	}
	app.TemplateCache = tc
	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.ListenAndServe(portNumber, nil)
}
