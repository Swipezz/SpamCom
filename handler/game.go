package handler

import (
	"html/template"
	"main/service"
	"net/http"
)

type handler struct {
	query service.Query
}

func NewHandler(query service.Query) *handler {
	return &handler{query}
}

func (h *handler) MainMenu(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/views/index.html"))
	tmpl.Execute(w, nil)
}

func (h *handler) Test(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/views/test.html"))
	tmpl.Execute(w, nil)
}

func (h *handler) StyleTest(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/views/styletest.html"))
	tmpl.Execute(w, nil)
}

func (h *handler) Trash(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/views/trash.html"))
	tmpl.Execute(w, nil)
}

func (h *handler) Image(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/views/image.html"))
	tmpl.Execute(w, nil)
}
