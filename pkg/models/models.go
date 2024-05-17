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
		GetAll(title string, genres []string, filters Filters) ([]*Book, Metadata, error)
	}
	Mangas interface {
		Insert(manga *Manga) error
		Get(id int64) (*Manga, error)
		Update(manga *Manga) error
		Delete(id int64) error
		GetAll(title string, genres []string, filters Filters) ([]*Manga, Metadata, error)
	}
	Authors interface {
		Insert(author *Author) error
		Get(id int64) (*Author, error)
		Update(author *Author) error
		Delete(id int64) error
		GetAll(name string, id int64, filters Filters) ([]*Author, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Books:   BookModel{DB: db},
		Mangas:  MangaModel{DB: db},
		Authors: AuthorModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Books:   MockBookModel{},
		Authors: MockAuthorModel{},
	}
}
