package handler

import (
    "net/http"
    "log"
    "html/template"
)

func Home(w http.ResponseWriter,r *http.Request) {
    tmpl, err := template.ParseFiles("templates/base.html")
    if err  != nil {
        // handle err
    }
    if r.URL.Path != "/" {
        http.NotFound(w,r)
        return
    }
    err2 := tmpl.ExecuteTemplate(w, "base", "Home Page!")
    if err2 != nil {
        log.Println(err2)
    }
}
