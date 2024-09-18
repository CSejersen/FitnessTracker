package storage

import (
	"github.com/Csejersen/fitnessTracker/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateExercise(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	store := NewSqliteExerciseStore(db)

	exercise := &models.Exercise{Name: "Push-up", UserID: 1}

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM Users WHERE ID=\\?\\)").
		WithArgs(exercise.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	mock.ExpectExec("INSERT INTO Exercises \\(name, userID\\) VALUES \\(\\?, \\?\\)").
		WithArgs(exercise.Name, exercise.UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = store.CreateExercise(exercise)
	assert.NoError(t, err)
}

func TestCreateExercise_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	store := NewSqliteExerciseStore(db)

	exercise := &models.Exercise{Name: "Push-up", UserID: 1}
	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM Users WHERE ID=\\?\\)").
		WithArgs(exercise.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	err = store.CreateExercise(exercise)
	assert.Error(t, err)
	assert.Equal(t, "User not found", err.Error())
}
