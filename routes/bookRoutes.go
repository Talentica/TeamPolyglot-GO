package routes

import (
	"net/http"

	"book-shelve/TeamPolyglot-GO/app/book"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//SetBookRoutes - set routes fot authentication
func SetBookRoutes(router *mux.Router, commonHandlers alice.Chain, webserviceHandler *book.WebServiceHandler) *mux.Router {
	router.Handle("/bookSharing/api/v1.0/books", commonHandlers.ThenFunc(func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.Get(res, req)
	})).Methods("GET")

	return router
}
