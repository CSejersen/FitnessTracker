package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Csejersen/fitnessTracker/handlers"
	"github.com/Csejersen/fitnessTracker/server"
	"github.com/Csejersen/fitnessTracker/storage"
)

func main() {
	db, err := sql.Open("sqlite3", "file:fitness_tracker.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}
	defer db.Close()

	err = storage.CreateSchema(db)
	if err != nil {
		log.Fatalf("could not create schema: %v", err)
	}

	exerciseStore := storage.NewSqliteExerciseStore(db)
	exerciseHandler := &handlers.ExerciseHandler{
		Store: exerciseStore,
	}

	userStore := storage.NewSqliteUserStore(db)
	userHandler := &handlers.UserHandler{
		Store: userStore,
	}

	server := server.NewAPIServer(":3000", exerciseHandler, userHandler)
	server.Run()
}
