package book

import (
	"book-shelve/TeamPolyglot-GO/interfaces"
	"fmt"
)

//DbBookRepo -
type DbBookRepo interfaces.DbRepo

//NewDbRepo -
func NewDbRepo(dbHandler interfaces.DbHandler) *DbBookRepo {
	dbBookRepo := new(DbBookRepo)
	dbBookRepo.DbHandler = dbHandler
	return dbBookRepo
}

//Get - return all books
func (repo *DbBookRepo) Get() []Book {
	var books []Book
	row := repo.DbHandler.Query("SELECT id, Name FROM Books")
	var book Book
	for row.Next() {
		row.Scan(&book.ID, &book.Name)
		books = append(books, book)
	}
	return books
}

//GetById - return book by Id
func (repo *DbBookRepo) GetById(id int) Book {
	var book Book
	row := repo.DbHandler.Query(fmt.Sprintf("SELECT id, Name FROM Books WHERE ID = %v", id))
	for row.Next() {
		row.Scan(&book.ID, &book.Name)
	}
	return book
}
