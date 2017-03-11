package controllers

import (
	"html/template"
	"net/http"

	"github.com/didiyudha/go-mvc-pattern/templates"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = templates.Template()
}

// Index home page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Register for new user
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "registers.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
