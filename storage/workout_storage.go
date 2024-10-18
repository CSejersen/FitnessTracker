package storage

import (
	"database/sql"
	"fmt"

	"github.com/Csejersen/fitnessTracker/models"
	_ "github.com/mattn/go-sqlite3"
)

type WorkoutStore interface {
	CreateWorkout(*models.Workout) error
	GetAllWorkoutsByUserID(int) ([]models.Workout, error)
	GetWorkoutByID(int) (*models.Workout, error)
	DeleteWorkoutByID(int) error
	AddExercise(exerciseID int, workoutID int) error
}

type SqliteWorkoutStore struct {
	DB *sql.DB
}

func NewSqliteWorkoutStore(db *sql.DB) *SqliteWorkoutStore {
	return &SqliteWorkoutStore{
		DB: db,
	}
}

func (s *SqliteWorkoutStore) CreateWorkout(workout *models.Workout) error {
	if _, err := s.checkUserExists(workout.UserID); err != nil {
		return err
	}

	query := "INSERT INTO Workouts (name, userID) VALUES (?, ?)"
	_, err := s.DB.Exec(query, workout.Name, workout.UserID)
	if err != nil {
		return fmt.Errorf("Failed to insert workout: %w", err)
	}
	return nil
}

func (s *SqliteWorkoutStore) GetAllWorkoutsByUserID(id int) ([]models.Workout, error) {
	if _, err := s.checkUserExists(id); err != nil {
		return nil, err
	}

	query := "SELECT ID, name, exercises FROM Workouts WHERE UserID=?"
	rows, err := s.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to query Workouts: %v", err)
	}

	defer rows.Close()

	var workouts []models.Workout
	for rows.Next() {
		workout := models.Workout{}
		if err := rows.Scan(&workout.ID, &workout.Name, &workout.Exercises); err != nil {
			return nil, fmt.Errorf("Failed to scan rows: %v", err)
		}
		workouts = append(workouts, workout)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %v", err)
	}
	return workouts, nil
}

func (s *SqliteWorkoutStore) GetWorkoutByID(id int) (*models.Workout, error) {
	query := "SELECT ID, userID, name, exercises FROM Workouts WHERE ID=?"
	workout := &models.Workout{}
	if err := s.DB.QueryRow(query, id).Scan(&workout.ID, &workout.UserID, &workout.Name, &workout.Exercises); err != nil {
		return nil, fmt.Errorf("Failed to scan row %v", err)
	}
	return workout, nil
}

func (s *SqliteWorkoutStore) DeleteWorkoutByID(id int) error {
	query := "DELETE FROM Workouts WHERE id=?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete workout %v", err)
	}
	return nil
}

func (s *SqliteWorkoutStore) AddExercise(exerciseID int, workoutID int) error {
	query := "INSERT INTO WorkoutExercises (exerciseID, workoutID) VALUES (?, ?)"
	_, err := s.DB.Exec(query, exerciseID, workoutID)
	if err != nil {
		return fmt.Errorf("Failed to add exercise %v", err)
	}
	return nil
}

func (s *SqliteWorkoutStore) checkUserExists(id int) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM Users WHERE ID=?)"
	var exists bool
	if err := s.DB.QueryRow(query, id).Scan(&exists); err != nil {
		return false, fmt.Errorf("Failed to check if user exists: %w", err)
	}
	if !exists {
		return false, fmt.Errorf("User not found")
	}
	return true, nil
}
