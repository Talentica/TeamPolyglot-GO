package main

import (
	"book-shelve/TeamPolyglot-GO/app/book"
	"book-shelve/TeamPolyglot-GO/infrastructure"
	"book-shelve/TeamPolyglot-GO/middlewares"
	"book-shelve/TeamPolyglot-GO/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func main() {
	dbhandler := infrastructure.NewMySQLHandler("sa:real$123@tcp(localhost:3306)/books")
	bookInteractor := new(book.Interactor)
	bookInteractor.Repository = book.NewDbRepo(dbhandler)

	bookWebServiceHandler := new(book.WebServiceHandler)
	bookWebServiceHandler.Interactor = *bookInteractor

	commonHandlers := alice.New(middlewares.LoggingHandler, middlewares.RecoverHandler)
	router := mux.NewRouter()
	routes.SetBookRoutes(router, commonHandlers, bookWebServiceHandler)
	http.ListenAndServe(":8080", router)
}
