package handler

import (
	"github.com/hackeraks/bookings/pkg/config"
	"github.com/hackeraks/bookings/pkg/model"
	"github.com/hackeraks/bookings/pkg/render"

	"net/http"
)

// hold data sent to template

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &model.TemplateData{})
}

// Resrvation renders reservation page
func (m *Repository) Resrvation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.html", &model.TemplateData{})
}

// Genrals renders genrals page
func (m *Repository) Genrals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "genrals.page.html", &model.TemplateData{})
}

// Majors renders majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.html", &model.TemplateData{})
}

// Availability renders majors page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availablity.page.html", &model.TemplateData{})
}

// Contact renders contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.html", &model.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, fmt.Sprintf("This is the about page and sum is %d", AddValue(2, 9)))
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}
