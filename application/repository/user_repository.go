package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewInstance(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(name, email, hashedPassword string, createdAt, updatedAt time.Time) (int64, error) {
	result, err := ur.db.Exec("INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		name, email, hashedPassword, createdAt, updatedAt)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
func (ur *UserRepository) FetchUser() (map[string]string, error) {
	var (
		id         string
		username   string
		email      string
		password   string
		created_at string
		updated_at string
	)

	err := ur.db.QueryRow("SELECT * FROM users").Scan(
		&id,
		&username,
		&email,
		&password,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	userData := map[string]string{
		"id":         id,
		"username":   username,
		"email":      email,
		"password":   password,
		"created_at": created_at,
		"updated_at": updated_at,
	}

	return userData, nil
}
func (ur *UserRepository) UpdateUser(id int, name, email, password string, updatedAt time.Time) error {
	result, err := ur.db.Exec("UPDATE users SET name=?, email=?, password=?, updated_at=? WHERE id=?",
		name, email, password, updatedAt, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		// Tidak ada baris yang terpengaruh, user dengan ID tersebut tidak ditemukan
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}
func (ur *UserRepository) DeleteUser(id int) error {
	result, err := ur.db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		// Tidak ada baris yang terpengaruh, user dengan ID tersebut tidak ditemukan
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}
func (ur *UserRepository) GetUser(id int) (map[string]string, error) {
	var (
		username   string
		email      string
		password   string
		created_at string
		updated_at string
	)

	err := ur.db.QueryRow("SELECT * FROM users where id=?", id).Scan(
		&id,
		&username,
		&email,
		&password,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	userData := map[string]string{
		"username":   username,
		"email":      email,
		"password":   password,
		"created_at": created_at,
		"updated_at": updated_at,
	}

	return userData, nil
}
