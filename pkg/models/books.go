package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"library-app/pkg/validator"
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	AuthorId  int64     `json:"author_id,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateBook(v *validator.Validator, book *Book) {
	v.Check(book.Title != "", "title", "must be provided")
	v.Check(len(book.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(book.Year != 0, "year", "must be provided")
	v.Check(book.Year >= 1888, "year", "must be greater than 1888")
	v.Check(book.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(book.Genres != nil, "genres", "must be provided")
	v.Check(len(book.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(book.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(book.Genres), "genres", "must not contain duplicate values")
}

type BookModel struct {
	DB *sql.DB
}

func (m BookModel) Insert(book *Book) error {

	query := `
        INSERT INTO books (title, year, author_id, genres)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version`

	args := []interface{}{book.Title, book.Year, book.AuthorId, pq.Array(book.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
}

func (m BookModel) Get(id int64) (*Book, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
        SELECT id, created_at, title, year, author_id, genres, version
        FROM books
        WHERE id = $1`
	var book Book
	err := m.DB.QueryRow(query, id).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.Title,
		&book.Year,
		&book.AuthorId,
		pq.Array(&book.Genres),
		&book.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &book, nil
}

func (m BookModel) Update(book *Book) error {
	query := `
        UPDATE books 
        SET title = $1, year = $2, author_id = $3, genres = $4, version = version + 1
        WHERE id = $5
        RETURNING version`
	args := []interface{}{
		book.Title,
		book.Year,
		book.AuthorId,
		pq.Array(book.Genres),
		book.ID,
	}
	return m.DB.QueryRow(query, args...).Scan(&book.Version)
}

func (m BookModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
        DELETE FROM books
        WHERE id = $1`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (m BookModel) GetAll(title string, genres []string, filters Filters) ([]*Book, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), id, created_at, title, year, author_id, genres, version
        FROM books
        WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
        AND (genres @> $2 OR $2 = '{}')     
        ORDER BY %s %s, id ASC
        LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{title, pq.Array(genres), filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	books := []*Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(
			&totalRecords,
			&book.ID,
			&book.CreatedAt,
			&book.Title,
			&book.Year,
			&book.AuthorId,
			pq.Array(&book.Genres),
			&book.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		books = append(books, &book)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return books, metadata, nil
}

type MockBookModel struct{}

func (m MockBookModel) Insert(book *Book) error {
	// Мокируем действие...
	return nil
}

func (m MockBookModel) Get(id int64) (*Book, error) {
	// Мокируем действие...
	return nil, nil
}

func (m MockBookModel) Update(book *Book) error {
	// Мокируем действие...
	return nil
}

func (m MockBookModel) Delete(id int64) error {
	// Мокируем действие...
	return nil
}

func (m MockBookModel) GetAll(title string, genres []string, filters Filters) ([]*Book, Metadata, error) {
	return nil, Metadata{}, nil
}
