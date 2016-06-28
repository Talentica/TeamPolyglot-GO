package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vinod4006/TeamPolyglot-GO/assignments/ToDos/infrastructure"
	"github.com/vinod4006/TeamPolyglot-GO/assignments/ToDos/interfaces"
	"github.com/vinod4006/TeamPolyglot-GO/assignments/ToDos/usecases"
)

func main() {

	dbhandler := infrastructure.NewMysqlHandler("sa:real$123@tcp(localhost:3306)/todos")
	taskInteractor := new(usecases.TaskInteractor)
	taskInteractor.TaskRepository = interfaces.NewDbTaskRepo(dbhandler)

	webServiceHandler := interfaces.WebserviceHandler{}
	webServiceHandler.TaskInteractor = taskInteractor
	router := mux.NewRouter()
	router.HandleFunc("/todo/api/v1.0/tasks", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello req")
		webServiceHandler.Tasks(res, req)
	}).Methods("GET")
	http.ListenAndServe(":8080", router)
}
