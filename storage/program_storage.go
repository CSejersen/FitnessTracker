package storage

import (
	"database/sql"
	"fmt"

	"github.com/Csejersen/fitnessTracker/models"
	_ "github.com/mattn/go-sqlite3"
)

type ProgramStore interface {
	CreateProgram(*models.Program) error
	GetProgramByUserID(id int) ([]*models.Program, error)
	DeleteProgram(id int) error
	AddWorkout(workoutId int, programId int) error
}

type SqliteProgramStore struct {
	DB *sql.DB
}

func NewSqliteProgramStore(db *sql.DB) *SqliteProgramStore {
	return &SqliteProgramStore{
		DB: db,
	}
}

func (s *SqliteProgramStore) CreateProgram(program *models.Program) error {
	query := "INSERT INTO Programs (userID, name, split, perWeek) VALUES (?, ?, ?, ?)"
	_, err := s.DB.Exec(query, program.UserID, program.Name, program.Split, program.PerWeek)

	if err != nil {
		return fmt.Errorf("Failed to create program: %v", err)
	}

	return nil
}

func (s *SqliteProgramStore) GetProgramByUserID(id int) ([]*models.Program, error) {
	query := `SELECT ID, userID, name, split, perWeek from Programs WHERE userID=?`

	var programs []*models.Program

	rows, err := s.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to query database %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		program := &models.Program{}

		if err := rows.Scan(&program.ID, &program.UserID, &program.Name, &program.Split, &program.PerWeek); err != nil {
			return nil, fmt.Errorf("Failed to scan row %v", err)
		}

		programs = append(programs, program)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %v", err)
	}

	return programs, nil
}

func (s *SqliteProgramStore) DeleteProgram(id int) error {
	query := "DELETE FROM Programs WHERE id=?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete program %v", err)
	}
	return nil
}

func (s *SqliteProgramStore) AddWorkout(workoutId int, programId int) error {
	query := "INSERT INTO ProgramWorkouts (workoutID, programID) VALUES (?, ?)"

	_, err := s.DB.Exec(query, workoutId, programId)
	if err != nil {
		return fmt.Errorf("Failed to add workout: %v", err)
	}

	return nil
}
