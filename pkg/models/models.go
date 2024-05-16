package models

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Books interface {
		Insert(book *Book) error
		Get(id int64) (*Book, error)
		Update(book *Book) error
		Delete(id int64) error
		GetAll(title string, genres []string, filters Filters) ([]*Book, error)
	}
	//Author interface {
	//	Insert(book *Book) error
	//	Get(id int64) (*Book, error)
	//	Update(book *Book) error
	//	Delete(id int64) error
	//	GetAll(title string) ([]*Author, error)
	//}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Books: BookModel{DB: db},
		//Author: AuthorModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Books: MockBookModel{},
		//Author: MockAuthorModel{},
	}
}
