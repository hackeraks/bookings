package handler

import (
	"encoding/json"
	"fmt"
	"log"

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
	render.RenderTemplate(w, r, "home.page.html", &model.TemplateData{})
}

// Resrvation renders reservation page
func (m *Repository) Resrvation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &model.TemplateData{})
}

// Genrals renders genrals page
func (m *Repository) Genrals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "genrals.page.html", &model.TemplateData{})
}

// Majors renders majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.html", &model.TemplateData{})
}

// Availability renders majors page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availablity.page.html", &model.TemplateData{})
}

// Postvailability renders majors page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))

	// render.RenderTemplate(w, "search-availablity.page.html", &model.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"Ok"`
	Message string `json:"message"`
}

// PostAvailabilityjson send json response
func (m *Repository) PostAvailabilityJson(w http.ResponseWriter, r *http.Request) {
	jout := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(jout, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &model.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, fmt.Sprintf("This is the about page and sum is %d", AddValue(2, 9)))
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, r, "about.page.html", &model.TemplateData{
		StringMap: stringMap,
	})
}
