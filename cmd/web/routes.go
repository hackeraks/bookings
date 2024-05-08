package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/hackeraks/bookings/pkg/config"
	handler "github.com/hackeraks/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	mux.Get("/generals-quarters", handler.Repo.Genrals)
	mux.Get("/majors-suite", handler.Repo.Majors)
	mux.Get("/search-availability", handler.Repo.Availability)
	mux.Get("/contact", handler.Repo.Contact)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
