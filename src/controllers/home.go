package controllers

import (
	"html/template"
	"net/http"
)

type HomeController struct {
	http.Handler
	Template *template.Template
}

func (this *HomeController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	this.Template.Execute(w, nil)
}
