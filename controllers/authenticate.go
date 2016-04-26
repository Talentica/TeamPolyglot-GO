package controllers

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
