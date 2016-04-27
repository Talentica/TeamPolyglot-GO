package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Talentica/TeamPolyglot-GO/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

//Login - Authenticates requested user, with valid credentials
func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	response := new(models.Response)
	if err := decoder.Decode(&requestUser); err == nil {

		token := jwt.New(jwt.SigningMethodHS256)

		token.Claims["name"] = "requestUser.Username"
		token.Claims["iat"] = time.Now().Unix()
		token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		jwtString, errJwt := token.SignedString([]byte("secret"))

		if errJwt == nil {

			response.Data = jwtString
			response.Message = "success"
			response.Status = true
			w.WriteHeader(http.StatusOK)

		} else {
			response.Data = errJwt.Error()
			response.Message = "failure"
			response.Status = false
			w.WriteHeader(http.StatusInternalServerError)
		}

	} else {
		response.Data = err.Error()
		response.Message = "failure"
		response.Status = false
		w.WriteHeader(http.StatusBadRequest)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
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
