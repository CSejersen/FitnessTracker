package main

import (
	"log"

	"github.com/Csejersen/fitnessTracker/handlers"
	"github.com/Csejersen/fitnessTracker/models"
	"github.com/Csejersen/fitnessTracker/server"
	"github.com/Csejersen/fitnessTracker/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dbFile := "fitness_tracker.db"

	// Open a connection to the SQLite database
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&models.Exercise{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	exerciseStore := storage.NewGormExerciseStore(db)

	exerciseHandler := &handlers.ExerciseHandler{
		Store: exerciseStore,
	}
	server := server.NewAPIServer(":3000", exerciseHandler)
	server.Run()
}
