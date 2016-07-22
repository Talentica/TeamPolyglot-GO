package book

import (
	"encoding/json"
	"net/http"
)

//WebServiceHandler - s
type WebServiceHandler struct {
	Interactor Interactor
}

//Get - web service handler
func (handler WebServiceHandler) Get(res http.ResponseWriter, req *http.Request) {
	books := handler.Interactor.Get()
	json.NewEncoder(res).Encode(books)
}

//GetByID - get book by id
func (handler WebServiceHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	id := 1
	book := handler.Interactor.GetById(id)
	json.NewEncoder(res).Encode(book)
}
