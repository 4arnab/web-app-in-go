package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RENDERING A HTML FILES

func RenderTemplates(w http.ResponseWriter, temp string) {
	parsedTemplate, _ := template.ParseFiles("../../templates/" + temp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates", err)
		return
	}
}
