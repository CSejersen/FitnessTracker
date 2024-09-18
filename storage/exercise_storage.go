package storage

import (
	"database/sql"
	"fmt"

	"github.com/Csejersen/fitnessTracker/models"
	_ "github.com/mattn/go-sqlite3"
)

type ExerciseStore interface {
	CreateExercise(*models.Exercise) error
	GetAllExercises() ([]models.Exercise, error)
	GetExercisesByUserID(int) ([]*models.Exercise, error)
	DeleteExerciseByID(int) error
}

type SqliteExerciseStore struct {
	DB *sql.DB
}

func NewSqliteExerciseStore(db *sql.DB) *SqliteExerciseStore {
	return &SqliteExerciseStore{
		DB: db,
	}
}

func (s *SqliteExerciseStore) CreateExercise(exercise *models.Exercise) error {
	// Check if the user exists
	query := "SELECT EXISTS(SELECT 1 FROM Users WHERE ID=?)"
	var exists bool
	if err := s.DB.QueryRow(query, exercise.UserID).Scan(&exists); err != nil {
		return fmt.Errorf("Failed to check if user exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("User not found")
	}

	query = "INSERT INTO Exercises (name, userID) VALUES (?, ?)"
	_, err := s.DB.Exec(query, exercise.Name, exercise.UserID)
	if err != nil {
		return fmt.Errorf("Failed to insert exercise: %w", err)
	}
	return nil
}

func (s *SqliteExerciseStore) GetExercisesByUserID(id int) ([]*models.Exercise, error) {
	query := `SELECT ID, userID, name from Exercises WHERE userID=?`
	var exercises []*models.Exercise

	rows, err := s.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to query database %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		exercise := &models.Exercise{}

		if err := rows.Scan(&exercise.ID, &exercise.UserID, &exercise.Name); err != nil {
			return nil, fmt.Errorf("Failed to scan row %v", err)
		}

		exercises = append(exercises, exercise)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %v", err)
	}
	return exercises, nil
}

func (s *SqliteExerciseStore) GetAllExercises() ([]models.Exercise, error) {
	query := `SELECT ID, userID, name from Exercises`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get exercises: %w", err)
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var exercise models.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.UserID, &exercise.Name); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		exercises = append(exercises, exercise)

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error occurred while iterating rows: %w", err)
		}
	}
	return exercises, nil
}

func (s *SqliteExerciseStore) DeleteExerciseByID(id int) error {
	query := "DELETE FROM Exercises WHERE id=?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete exercise %v", err)
	}
	return nil
}
