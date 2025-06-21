package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.html")
}

// RENDERING A HTML FILES
func renderTemplates(w http.ResponseWriter, temp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + temp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates" + err.Error())
		return
	}
}
