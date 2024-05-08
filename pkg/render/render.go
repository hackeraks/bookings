package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/hackeraks/bookings/pkg/config"
	"github.com/hackeraks/bookings/pkg/model"
)

var tc = make(map[string]*template.Template)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *model.TemplateData) *model.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, td *model.TemplateData) {
	// get the tempalte cache
	tc := app.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn't get template")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, er := buf.WriteTo(w)
	if er != nil {
		fmt.Println(er)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.html")
	log.Println(pages)
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil
}
