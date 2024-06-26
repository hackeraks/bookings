package main

import (
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/hackeraks/bookings/internal/config"
	handler "github.com/hackeraks/bookings/internal/handlers"
	"github.com/hackeraks/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
	}
	app.TemplateCache = tc
	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)
	render.NewTemplate(&app)
	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)

	fmt.Println(fmt.Sprintf("Starting service on port %s", portNumber))

	// _ = http.ListenAndServe(":8080", nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
