package usecases

import "github.com/vinod4006/TeamPolyglot-GO/assignments/ToDos/domain"

//TaskInteractor - use case handler
type TaskInteractor struct {
	TaskRepository domain.TaskRepository
}

//GetTasks - all tasks
func (interactor *TaskInteractor) GetTasks() []domain.Task {
	var tasks []domain.Task
	tasks = interactor.TaskRepository.GetAllTasks()
	return tasks
}
