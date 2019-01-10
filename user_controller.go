package main

import (
	"encoding/json"
	"fmt"
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
	resp, err := http.Get(usersURL + "?q=" + name)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}

	var result UsersSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
		resp.Body.Close()
		return
	}

	templates := template.Must(template.ParseFiles("views/result.html"))

	if err := templates.ExecuteTemplate(w, "result.html", result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
