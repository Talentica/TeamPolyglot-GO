package routers

import (
	"github.com/Talentica/TeamPolyglot-GO/controllers"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/authenticate/login", controllers.Login).Methods("POST")
	router.HandleFunc("/authenticate/logout", controllers.Logout).Methods("POST")

	return router
}
