package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Csejersen/fitnessTracker/config"
	database "github.com/Csejersen/fitnessTracker/db"
	"github.com/Csejersen/fitnessTracker/handlers"
	"github.com/Csejersen/fitnessTracker/server"
	"github.com/Csejersen/fitnessTracker/storage"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "file:fitness_tracker.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}
	defer db.Close()

	err = database.CreateSchema(db)
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

	loginHandler := &handlers.LoginHandler{
		Store: userStore,
		Cfg:   *cfg,
	}

	workoutStore := storage.NewSqliteWorkoutStore(db)
	workoutHandler := &handlers.WorkoutHandler{
		Store: workoutStore,
	}

	programStore := storage.NewSqliteProgramStore(db)
	programHandler := &handlers.ProgramHandler{
		Store: programStore,
	}

	server := server.NewAPIServer(cfg.Port, exerciseHandler, userHandler, loginHandler, workoutHandler, programHandler)
	err = server.Run()
	fmt.Println(err)
}
