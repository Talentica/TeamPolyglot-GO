package interfaces

import (
	"fmt"

	"github.com/vinod4006/TeamPolyglot-GO/assignments/ToDos/domain"
)

//DbHandler -
type DbHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

//Row
type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

//DbRepo -
type DbRepo struct {
	dbHandler DbHandler
}

//DbTaskRepo -
type DbTaskRepo DbRepo

//NewDbTaskRepo -
func NewDbTaskRepo(dbHandler DbHandler) *DbTaskRepo {
	dbTaskRepo := new(DbTaskRepo)

	dbTaskRepo.dbHandler = dbHandler
	return dbTaskRepo
}

//Store -
func (repo *DbTaskRepo) Store(task domain.Task) {
	repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO Tasks (id, title, descripion, done) VALUES ('%d', '%v', '%v', '%d')", task.ID, task.Title, task.Desc, task.Done))
}

//FindById -
func (repo *DbTaskRepo) FindById(id int) domain.Task {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT title, descripion, done FROM Tasks WHERE id = '%d' LIMIT 1", id))
	var desc string
	var title string
	var done int
	row.Next()
	row.Scan(&title, &desc, &done)
	task := domain.Task{ID: id, Title: title, Desc: desc, Done: done}
	return task
}

//GetAllTasks() -
func (repo *DbTaskRepo) GetAllTasks() []domain.Task {
	var tasks []domain.Task
	row := repo.dbHandler.Query("SELECT id, title, descripion, done FROM Tasks")
	var task domain.Task
	for row.Next() {
		row.Scan(&task.ID, &task.Title, &task.Desc, &task.Done)
		tasks = append(tasks, task)
	}
	return tasks
}
