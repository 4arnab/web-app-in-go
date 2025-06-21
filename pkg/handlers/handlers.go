package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.html")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.html")
}
