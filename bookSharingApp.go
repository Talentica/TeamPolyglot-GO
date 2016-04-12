package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func main() {
	router.HandleFunc("/", indexPageHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
