package book

//IRepository -
type IRepository interface {
	Get() []Book
	GetById(id int) Book
}

//Book - domain entity
type Book struct {
	ID   int
	Name string
}
