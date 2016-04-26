package routers

import (
	"github.com/Talentica/TeamPolyglot-GO/controllers"
	"github.com/gorilla/mux"
)

//SetAuthenticationRoutes - set routes fot authentication
func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/authenticate/login", controllers.Login).Methods("POST")
	router.HandleFunc("/authenticate/logout", controllers.Logout).Methods("POST")
	router.HandleFunc("/authenticate/revoke-token/{token}", controllers.RevokeToken).Methods("PUT")
	router.HandleFunc("/authenticate/forgot-password/{username}", controllers.ForgotPassword).Methods("POST")
	router.HandleFunc("/authenticate/reset-password/{token}", controllers.ResetPassword).Methods("GET")
	return router
}
