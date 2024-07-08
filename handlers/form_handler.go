package handler

import (
	"net/http"
	"html/template"
)

func Form(w http.ResponseWriter,r *http.Request) {
	tmpl, err := template.ParseFiles("templates/includes/form.html")
	if err != nil {
		// handle error
	}
	err = tmpl.ExecuteTemplate(w, "form", "hello form!")
    if err != nil {
        // handle error
    }
}

