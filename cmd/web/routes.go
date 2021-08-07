package main

import (
	"fmt"
	"hotelsystem/pkg/config"
	"hotelsystem/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	fileServer := http.FileServer(http.Dir("./static/"))
	fileDir := http.Dir("./static")
	fmt.Println("file dir is", fileDir)
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	return mux

}
