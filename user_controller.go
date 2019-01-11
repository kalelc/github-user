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

func Result(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	var result GithubSearchEngine
	err := result.Search(name)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templates := template.Must(template.ParseFiles("views/result.html"))

	if err := templates.ExecuteTemplate(w, "result.html", result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
