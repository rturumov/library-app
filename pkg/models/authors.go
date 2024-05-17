package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Author struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type AuthorModel struct {
	DB *sql.DB
}

func (m AuthorModel) Insert(author *Author) error {

	query := `
        INSERT INTO authors (name)
        VALUES ($1)
        RETURNING id`

	args := []interface{}{author.Name}

	return m.DB.QueryRow(query, args...).Scan(&author.Id)
}

func (m AuthorModel) Get(id int64) (*Author, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
        SELECT id, name
        FROM authors
        WHERE id = $1`
	var author Author
	err := m.DB.QueryRow(query, id).Scan(
		&author.Id,
		&author.Name,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &author, nil
}

func (m AuthorModel) Update(author *Author) error {
	query := `
        UPDATE authors
        SET name = $1
        WHERE id = $2`

	_, err := m.DB.Exec(query, author.Name, author.Id)
	if err != nil {
		return err // Return the error if any occurred during the execution
	}

	return nil
}

func (m AuthorModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
        DELETE FROM authors
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

func (m AuthorModel) GetAll(Name string, id int64, filters Filters) ([]*Author, error) {
	query := `
        SELECT id, name
        FROM authors
        WHERE (LOWER(name) = LOWER($1) OR $1 = '')      
        ORDER BY id`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// Pass the title and genres as the placeholder parameter values.
	rows, err := m.DB.QueryContext(ctx, query, Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	authors := []*Author{}
	for rows.Next() {
		var author Author
		err := rows.Scan(
			&author.Id,
			&author.Name,
		)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

type MockAuthorModel struct{}

func (m MockAuthorModel) Insert(author *Author) error {
	// Мокируем действие...
	return nil
}

func (m MockAuthorModel) Get(id int64) (*Author, error) {
	// Мокируем действие...
	return nil, nil
}

func (m MockAuthorModel) Update(author *Author) error {
	// Мокируем действие...
	return nil
}

func (m MockAuthorModel) Delete(id int64) error {
	// Мокируем действие...
	return nil
}

func (m MockAuthorModel) GetAll(Name string, id int64, filters Filters) ([]*Author, error) {
	return nil, nil
}
