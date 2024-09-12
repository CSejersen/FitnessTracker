package storage

import (
	"database/sql"

	"github.com/Csejersen/fitnessTracker/models"
	_ "github.com/mattn/go-sqlite3"
)

type WorkoutStore interface {
	CreateWorkout(*models.Exercise) error
	GetAllWorkouts() ([]models.Exercise, error)
	GetWorkoutByID(int) (*models.Exercise, error)
	DeleteWorkoutByID(int) error
}

type SqliteWorkoutStore struct {
	DB *sql.DB
}

func NewSqliteWorkoutStore(db *sql.DB) *SqliteWorkoutStore {
	return &SqliteWorkoutStore{
		DB: db,
	}
}

func (s *SqliteExerciseStore) CreateWorkout(workout *models.Workout) error {
	return nil
}
