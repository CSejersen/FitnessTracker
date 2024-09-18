package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSchema(db *sql.DB) error {
	schema := `
		CREATE TABLE IF NOT EXISTS Users (
				ID INTEGER PRIMARY KEY AUTOINCREMENT,
				username TEXT NOT NULL UNIQUE,
				encryptedPassword TEXT NOT NULL
		);

    CREATE TABLE IF NOT EXISTS Exercises (
        ID INTEGER PRIMARY KEY AUTOINCREMENT, 
        name TEXT NOT NULL,
				userID INTEGER NOT NULL,
				FOREIGN KEY (userID) REFERENCES Users(ID)
    );
		
		CREATE TABLE IF NOT EXISTS Workouts (
				ID INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL,
				userID INTEGER NOT NULL,
				FOREIGN KEY (userID) REFERENCES Users(ID)
		);
		
		CREATE TABLE IF NOT EXISTS ExercisesWorkouts(
				exerciseID INTEGER NOT NULL,
				workoutID INTEGER NOT NULL,
				FOREIGN KEY (exerciseID) REFERENCES Exercises(ID),
				FOREIGN KEY (workoutID) REFERENCES Workouts(ID),
				PRIMARY KEY (exerciseID, workoutID)
		);
    `
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	return nil
}
