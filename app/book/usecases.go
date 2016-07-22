package book

//Interactor - use cases for book
type Interactor struct {
	Repository IRepository
}

//Get - return all books
func (interactor *Interactor) Get() []Book {
	books := interactor.Repository.Get()
	return books
}

//GetById - return book by id
func (interactor *Interactor) GetById(id int) Book {
	book := interactor.Repository.GetById(id)
	return book
}
