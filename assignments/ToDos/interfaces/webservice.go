package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/vinod4006/TeamPolyglot-GO/assignments/ToDos/domain"
)

type TaskInteractor interface {
	GetTasks() []domain.Task
}

type WebserviceHandler struct {
	TaskInteractor TaskInteractor
}

func (handler WebserviceHandler) Tasks(res http.ResponseWriter, req *http.Request) {

	tasks := handler.TaskInteractor.GetTasks()
	json.NewEncoder(res).Encode(tasks)
}
