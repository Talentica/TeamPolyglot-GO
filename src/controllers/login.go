package controllers

import (
	//"html/template"
	"fmt"
	"net/http"
)

type LoginController struct {
	http.Handler
}

func (this *LoginController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/plain")
	email := req.FormValue("email")
	//password := req.FormValue("password")
	fmt.Fprintf(w, "Hello %s", email)
}
