package domain

//TaskRepository interface
type TaskRepository interface {
	Store(task Task)
	FindById(id int) Task
	GetAllTasks() []Task
}

//Task domain entity
type Task struct {
	ID    int
	Title string
	Desc  string
	Done  int
}
