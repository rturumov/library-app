package repository

type Authorization interface {
}

type BookList interface {
}

type Book interface {
}

type Repository struct {
	Authorization
	BookList
	Book
}

func NewRepository() *Repository {
	return &Repository{}
}
