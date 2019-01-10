package main

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates := template.Must(template.ParseFiles("views/index.html"))

		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Results(w http.ResponseWriter, r *http.Request) {
	r.FormValue("name")
}
