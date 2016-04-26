package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Talentica/TeamPolyglot-GO/models"
	"github.com/gorilla/mux"
)

//Login - Authenticates requested user, with valid credentials
func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	s := "Hello " + requestUser.Username
	w.Write([]byte(s))
	w.WriteHeader(http.StatusOK)
}

//Logout current logged-in user
func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}

//RevokeToken - Revokes the requested token
func RevokeToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	w.Write([]byte(token))
}

//ForgotPassword - Sends an email with the link to reset password
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	w.Write([]byte(username))
}

//ResetPassword - Randomely generates the password and sets the password for user identified by token
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	w.Write([]byte(token))
}
