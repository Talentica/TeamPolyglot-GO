package main

import (
	//"fmt"
	"net/http"
)

func main() {
	http.Handle("/", new(IndexPageHandler))
	http.ListenAndServe(":8081", nil)
}

type IndexPageHandler struct {
	http.Handler
}

func (this *IndexPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
