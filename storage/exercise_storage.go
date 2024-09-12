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
	GetExerciseByID(int) (*models.Exercise, error)
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
	query := "INSERT INTO Exercises (name) VALUES (?)"
	_, err := s.DB.Exec(query, exercise.Name)
	if err != nil {
		return fmt.Errorf("Failed to insert exercise: %w", err)
	}
	return nil
}

func (s *SqliteExerciseStore) GetExerciseByID(id int) (*models.Exercise, error) {
	query := `SELECT ID, name from Exercises WHERE ID=?`
	var exercise models.Exercise

	err := s.DB.QueryRow(query, id).Scan(&exercise.ID, &exercise.Name)
	if err != nil {
		return nil, fmt.Errorf("Failed to scan row int exercise: %v", err)
	}

	return &exercise, nil
}

func (s *SqliteExerciseStore) GetAllExercises() ([]models.Exercise, error) {
	query := `SELECT ID, name from Exercises`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get exercises: %w", err)
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var exercise models.Exercise
		if err := rows.Scan(&exercise.ID, &exercise.Name); err != nil {
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
		return fmt.Errorf("Failed to delete exercise with id %s: %v", id, err)
	}
	return nil
}
