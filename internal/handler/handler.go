package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

type Video struct {
	Url    string `json:"url"`
	Format string `json:"format"`
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/index.html")
}

func (v *Video) Download(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	format := r.FormValue("format")
}

func renderTemplate(w http.ResponseWriter, file string) {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		fmt.Println("Failed to parse html template", err)
	}
	tmpl.Execute(w, nil)
}
