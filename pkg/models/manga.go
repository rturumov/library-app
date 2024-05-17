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

type Manga struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	AuthorId  int64     `json:"author,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateManga(v *validator.Validator, manga *Manga) {
	v.Check(manga.Title != "", "title", "must be provided")
	v.Check(len(manga.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(manga.Year != 0, "year", "must be provided")
	v.Check(manga.Year >= 1888, "year", "must be greater than 1888")
	v.Check(manga.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(manga.Genres != nil, "genres", "must be provided")
	v.Check(len(manga.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(manga.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(manga.Genres), "genres", "must not contain duplicate values")
}

type MangaModel struct {
	DB *sql.DB
}

func (m MangaModel) Insert(manga *Manga) error {

	query := `
        INSERT INTO mangas (title, year, author_id, genres)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version`

	args := []interface{}{manga.Title, manga.Year, manga.AuthorId, pq.Array(manga.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&manga.ID, &manga.CreatedAt, &manga.Version)
}

func (m MangaModel) Get(id int64) (*Manga, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
        SELECT id, created_at, title, year, author_id, genres, version
        FROM mangas
        WHERE id = $1`
	var manga Manga
	err := m.DB.QueryRow(query, id).Scan(
		&manga.ID,
		&manga.CreatedAt,
		&manga.Title,
		&manga.Year,
		&manga.AuthorId,
		pq.Array(&manga.Genres),
		&manga.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &manga, nil
}

func (m MangaModel) Update(manga *Manga) error {
	query := `
        UPDATE mangas
        SET title = $1, year = $2, author_id = $3, genres = $4, version = version + 1
        WHERE id = $5
        RETURNING version`
	args := []interface{}{
		manga.Title,
		manga.Year,
		manga.AuthorId,
		pq.Array(manga.Genres),
		manga.ID,
	}
	return m.DB.QueryRow(query, args...).Scan(&manga.Version)
}

func (m MangaModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
        DELETE FROM mangas
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

func (m MangaModel) GetAll(title string, genres []string, filters Filters) ([]*Manga, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), id, created_at, title, year, author_id, genres, version
        FROM mangas
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
	mangas := []*Manga{}
	for rows.Next() {
		var manga Manga
		err := rows.Scan(
			&totalRecords,
			&manga.ID,
			&manga.CreatedAt,
			&manga.Title,
			&manga.Year,
			&manga.AuthorId,
			pq.Array(&manga.Genres),
			&manga.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		mangas = append(mangas, &manga)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return mangas, metadata, nil
}

type MockMangaModel struct{}

func (m MockMangaModel) Insert(manga *Manga) error {
	// Мокируем действие...
	return nil
}

func (m MockMangaModel) Get(id int64) (*Manga, error) {
	// Мокируем действие...
	return nil, nil
}

func (m MockMangaModel) Update(manga *Manga) error {
	// Мокируем действие...
	return nil
}

func (m MockMangaModel) Delete(id int64) error {
	// Мокируем действие...
	return nil
}

func (m MockMangaModel) GetAll(title string, genres []string, filters Filters) ([]*Manga, Metadata, error) {
	return nil, Metadata{}, nil
}
