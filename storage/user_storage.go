package storage

import (
	"database/sql"
	"fmt"

	"github.com/Csejersen/fitnessTracker/models"
	_ "github.com/mattn/go-sqlite3"
)

type UserStore interface {
	CreateUser(*models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(int) (*models.User, error)
	DeleteUserByID(int) error
}

type SqliteUserStore struct {
	DB *sql.DB
}

func NewSqliteUserStore(db *sql.DB) *SqliteUserStore {
	return &SqliteUserStore{
		DB: db,
	}
}

func (s *SqliteUserStore) CreateUser(user *models.User) error {
	query := "INSERT INTO Users (username, password) VALUES (?, ?)"
	_, err := s.DB.Exec(query, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}
	return nil
}

func (s *SqliteUserStore) GetUserByID(id int) (*models.User, error) {
	query := `SELECT ID, username from Users WHERE ID=?`
	var user models.User

	err := s.DB.QueryRow(query, id).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, fmt.Errorf("Failed to scan row: %v", err)
	}

	return &user, nil
}

func (s *SqliteUserStore) GetAllUsers() ([]models.User, error) {
	query := `SELECT ID, username from Users`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		users = append(users, user)

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error occurred while iterating rows: %w", err)
		}
	}
	return users, nil
}

func (s *SqliteUserStore) DeleteUserByID(id int) error {
	query := "DELETE FROM Users WHERE id=?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete user with id %s: %v", id, err)
	}
	return nil
}
