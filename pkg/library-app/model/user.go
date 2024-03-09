package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type User struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
type UserModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (u UserModel) Insert(user *User) error {
	query := `
			INSERT INTO users (username, password)
			VALUES($1, $2)
			RETURNING id, createdAt, updatedAt
			`
	args := []interface{}{user.Username, user.Password}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	return u.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)
}

func (u UserModel) GetAll() ([]*User, error) {
	query := `
		SELECT id, createdAt, updatedAt, username, password
		FROM users
		ORDER BY id
	`

	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserModel) GetById(id int) (*User, error) {
	query := `
		SELECT id, createdAt, updatedAt, username, password
		FROM users
		WHERE id = $1
		`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := u.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserModel) Update(user *User) error {
	query := `
		UPDATE users
		SET username = $1, password = $2
		WHERE id = $3
		RETURNING updatedAt
		`

	args := []interface{}{user.Username, user.Password, user.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return u.DB.QueryRowContext(ctx, query, args...).Scan(&user.UpdatedAt)
}

func (u UserModel) Delete(id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
		`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	_, err := u.DB.ExecContext(ctx, query, id)

	return err
}
